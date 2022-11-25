package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"simple-api/models"
	"simple-api/responses"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

func CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: validationErr.Error(), Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Email:    user.Email,
			Password: user.Password,
			Token:    user.Token,
			Role:     user.Role,
		}

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		newUser.Password = string(hashedPassword)

		//validate if there is existing user
		if response, ok := newUser.Validate(ctx); !ok {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		//all validation passed, then insert new data to users collection
		result, err := newUser.Create(ctx, newUser)

		if err != nil {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		response := responses.BaseResponse{Status: http.StatusCreated, Message: "success", Data: result}
		json.NewEncoder(rw).Encode(response)
	}
}

func GetToken() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		user := &models.User{}

		err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
		if err != nil {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: err.Error(), Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		result, err := user.Find(ctx)
		if err != nil && err == mongo.ErrNoDocuments {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusNotFound)
			response := responses.BaseResponse{Status: http.StatusNotFound, Message: "User not found", Data: map[string]interface{}{}}
			json.NewEncoder(rw).Encode(response)
			return
		}

		// create JWT token
		threeMinute := (time.Hour / 60) * 3
		tk := &models.Token{UserId: result.Id, Email: result.Email, Role: result.Role, RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(threeMinute)),
		}}
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
		tokenString, _ := token.SignedString([]byte(os.Getenv("token_secret_key")))

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusNotFound)
		response := responses.BaseResponse{Status: http.StatusNotFound, Message: "User not found", Data: map[string]interface{}{"email": user.Email, "token": tokenString}}
		json.NewEncoder(rw).Encode(response)
	}
}

package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-api/configs"
	"simple-api/models"
	"simple-api/responses"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}}
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			rw.WriteHeader(http.StatusBadRequest)
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: validationErr.Error(), Data: map[string]interface{}{}}
			rw.Header().Add("Content-Type", "application/json")
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

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
			rw.Header().Add("Content-Type", "application/json")
			json.NewEncoder(rw).Encode(response)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		response := responses.BaseResponse{Status: http.StatusCreated, Message: "success", Data: result}
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(response)
	}
}

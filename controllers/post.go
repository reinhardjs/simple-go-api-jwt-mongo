package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"simple-api/configs"
	"simple-api/models"
	"simple-api/responses"
	"simple-api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var post models.Post

		rw.Header().Add("Content-Type", "application/json")

		//validate the request body
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"error": err.Error()}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validationErr := utils.Validate.Struct(&post); validationErr != nil {
			response := responses.BaseResponse{Status: http.StatusBadRequest, Message: validationErr.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		newPost := models.Post{
			Id:          primitive.NewObjectID(),
			Title:       post.Title,
			Description: post.Description,
		}

		//all validation passed, then insert new data to users collection
		var postsCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
		result, err := postsCollection.InsertOne(ctx, newPost)

		if err != nil {
			response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		response := responses.BaseResponse{Status: http.StatusCreated, Message: "success", Data: result}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
	}
}

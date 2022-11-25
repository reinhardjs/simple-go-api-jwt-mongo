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

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
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

func UpdatePost() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["userId"]

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)
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

		update := bson.M{"title": post.Title, "description": post.Description}

		//all validation passed, then insert new data to users collection
		var postsCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
		result, err := postsCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})

		if err != nil {
			response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		response := responses.BaseResponse{Status: http.StatusOK, Message: "success", Data: result}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
	}
}

func DeletePost() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["userId"]

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)
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

		//all validation passed, then insert new data to users collection
		var postsCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
		result, err := postsCollection.DeleteOne(ctx, bson.M{"_id": objId})

		if err != nil {
			response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		response := responses.BaseResponse{Status: http.StatusOK, Message: "success", Data: result}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
	}
}

func GetPost() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(r)
		userId := params["userId"]

		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		rw.Header().Add("Content-Type", "application/json")

		var post models.Post

		//all validation passed, then insert new data to users collection
		var postsCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
		err := postsCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&post)

		if err != nil {
			response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		response := responses.BaseResponse{Status: http.StatusOK, Message: "success", Data: post}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
	}
}

func GetAllPost() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		var posts []models.Post

		//all validation passed, then insert new data to users collection
		var postsCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")
		results, err := postsCollection.Find(ctx, bson.M{})

		if err != nil {
			response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
			rw.WriteHeader(response.Status)
			json.NewEncoder(rw).Encode(response)
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singlePost models.Post
			if err = results.Decode(&singlePost); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				response := responses.BaseResponse{Status: http.StatusInternalServerError, Message: err.Error(), Data: map[string]interface{}{}}
				json.NewEncoder(rw).Encode(response)
			}

			posts = append(posts, singlePost)
		}

		response := responses.BaseResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"results": posts}}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
	}
}

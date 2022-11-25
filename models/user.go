package models

import (
	"context"
	"net/http"
	"simple-api/configs"
	"simple-api/responses"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password string             `json:"password,omitempty"`
	Token    string             `json:"token,omitempty"`
	Role     string             `json:"role,omitempty"`
}

func (account *User) Validate(context context.Context) (responses.BaseResponse, bool) {

	if !strings.Contains(account.Email, "@") {
		return responses.BaseResponse{Status: http.StatusBadRequest, Message: "Email address is required", Data: map[string]interface{}{}}, false
	}

	if len(account.Password) < 6 {
		return responses.BaseResponse{Status: http.StatusBadRequest, Message: "Password address is required", Data: map[string]interface{}{}}, false
	}

	filter := bson.M{"email": account.Email}

	//Email must be unique
	var result User
	var usersCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
	err := usersCollection.FindOne(context, filter).Decode(&result)

	if err != nil && err != mongo.ErrNoDocuments {
		return responses.BaseResponse{Status: http.StatusBadRequest, Message: "Connection error. Please retry", Data: map[string]interface{}{}}, false
	}
	if result.Email != "" {
		return responses.BaseResponse{Status: http.StatusConflict, Message: "Email address already in use by another user", Data: map[string]interface{}{}}, false
	}

	return responses.BaseResponse{Status: http.StatusOK, Message: "Requirement passed", Data: map[string]interface{}{}}, true
}

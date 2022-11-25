package models

import (
	"simple-api/configs"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var PostsCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")

type Post struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
}

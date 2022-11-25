package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Title       string             `json:"title,omitempty"`
	Description string             `json:"description,omitempty"`
}

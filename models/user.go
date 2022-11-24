package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	email    string             `json:"email,omitempty"`
	password string             `json:"password,omitempty"`
	token    string             `json:"token,omitempty"`
	role     string             `json:"role,omitempty"`
}

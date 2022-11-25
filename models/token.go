package models

import (
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
JWT claims struct
*/

type Token struct {
	UserId primitive.ObjectID
	Email  string
	Role   string
	jwt.RegisteredClaims
}

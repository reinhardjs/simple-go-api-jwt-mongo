package models

import "github.com/golang-jwt/jwt/v4"

/*
JWT claims struct
*/
type Token struct {
	Id    uint
	Email string
	Role  string
	jwt.StandardClaims
}

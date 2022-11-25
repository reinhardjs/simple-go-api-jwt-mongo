package models

import "github.com/golang-jwt/jwt/v4"

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

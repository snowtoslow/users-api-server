package model

import "github.com/dgrijalva/jwt-go"

type AuthCustomClaims struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

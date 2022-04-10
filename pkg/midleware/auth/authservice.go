package auth

import "github.com/dgrijalva/jwt-go"

type IAuthService interface {
	GenerateToken(email, password string) string
	ValidateToken(token string) (*jwt.Token, error)
}

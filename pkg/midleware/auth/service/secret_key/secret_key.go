package secret_key

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"users-api-server/pkg/midleware/auth"
	"users-api-server/pkg/midleware/auth/model"
)

type jwtServicesWithSecret struct {
	secretKey string
	issure    string
}

func JWTAuthSecretService(secret, issuer string) auth.IAuthService {
	return &jwtServicesWithSecret{
		secretKey: secret,
		issure:    issuer,
	}
}

func (j jwtServicesWithSecret) GenerateToken(email, password string) string {
	claims := &model.AuthCustomClaims{
		Name:     email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    j.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j jwtServicesWithSecret) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(j.secretKey), nil
	})
}

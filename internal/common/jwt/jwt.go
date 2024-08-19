package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	secret = []byte("a_longer_secret_key_for_hs512_that_is_at_least_64_bytes_long")
)

type Claims struct {
	UserType string `json:"user_type"`
}

func GenerateJWT(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"user_type": claims.UserType,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secret)
}

func ValidateJWT(tokenString string) (Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userType := claims["user_type"].(string)
		return Claims{
			UserType: userType,
		}, nil
	} else {
		return Claims{}, err
	}
}

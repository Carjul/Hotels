package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var API_SECRET = os.Getenv("API_SECRET")

var jwtKey = []byte(API_SECRET)

func GenerateJWT(user any) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = user
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

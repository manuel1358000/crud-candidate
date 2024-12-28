package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
)
var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,                     
		"exp":     time.Now().Add(time.Hour).Unix(), 
		"iat":     time.Now().Unix(),
	}

	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

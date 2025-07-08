package utils;

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func GenerateToken(userId int, role string) (string, error) {

	expirationTime := time.Now().Add(15 * time.Minute)
	claims := jwt.MapClaims{
		"userId": userId,
		"role":   role,
		"iat":    time.Now().Unix(),
		"exp":    expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("APP_SECRET")
	return token.SignedString([]byte(secretKey))
}

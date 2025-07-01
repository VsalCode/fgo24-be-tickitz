package controllers

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func generateToken(userId int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"userId": userId,
		"iat":    time.Now().Unix(),
		"exp":    expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("APP_SECRET")
	return token.SignedString([]byte(secretKey))
}
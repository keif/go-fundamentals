package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "horsebatterystaple"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // use unix format, two hours expiration
	})

	return token.SignedString(secretKey)
}

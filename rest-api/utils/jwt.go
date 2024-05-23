package utils

import (
	"errors"
	"fmt"
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

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// data, ok
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // because the signing method used in generate us a version of HMAC
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("Token parse error.")
	}

	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return errors.New("Token is invalid.")
	}

	//parsedClaims, ok := parsedToken.Claims.(jwt.MapClaims)
	//if !ok {
	//	return errors.New("Invalid token claims.")
	//}

	// how these values can be extracted
	// ok, email :=
	//email := parsedClaims["email"].(string)
	//userId := parsedClaims["userId"].(int64)

	return nil
}

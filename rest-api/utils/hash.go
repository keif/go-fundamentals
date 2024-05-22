package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bcryptHash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bcryptHash), err
}

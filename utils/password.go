package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswordHash(providedPassword string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(providedPassword))
}

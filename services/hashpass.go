package services

import (
	"golang.org/x/crypto/bcrypt"
)

type user interface {
	password() string
}

func GeneratePassword(password string) (string, error) {
	hashpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashpass), nil
}

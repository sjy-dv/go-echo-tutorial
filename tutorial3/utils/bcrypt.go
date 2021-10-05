package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func BcryptHashSync(password string) ([]byte, error) {

	salt := 8
	return bcrypt.GenerateFromPassword([]byte(password), salt)
}

func BcryptCompareSync(hash, password string) error {
	errors := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(errors)
	return errors
}
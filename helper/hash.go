package helper

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)
func HashPassword(password string) string {
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password),4)
	if err != nil {
		log.Println(err)
	}
	return string(passwordHashed)
}

func ComparePassword(password string, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password)); err != nil {
		return false
	}
	return true
}
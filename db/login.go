package db

import (
	"github.com/AdolfoCunquero/twitter-go/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (models.User, bool) {
	usr, exists, _ := UserExists(email)
	if !exists {
		return usr, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(usr.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usr, false
	}

	return usr, true
}

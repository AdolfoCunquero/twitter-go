package jwt

import (
	"time"

	"github.com/AdolfoCunquero/twitter-go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func TokenGenerator(usr models.User) (string, error) {
	secretKey := []byte("MastersDelDesarrollo123**")

	payload := jwt.MapClaims{
		"email":     usr.Email,
		"firstName": usr.FirstName,
		"lastName":  usr.LastName,
		"birthDate": usr.BirthDate,
		"location":  usr.Location,
		"webSite":   usr.WebSite,
		"_id":       usr.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

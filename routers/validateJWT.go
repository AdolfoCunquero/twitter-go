package routers

import (
	"errors"
	"strings"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func ValidateJWT(token string) (*models.Claim, bool, string, error) {
	secretKey := []byte("MastersDelDesarrollo123**")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("el formato del token no es valido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err == nil {
		_, exists, _ := db.UserExists(claims.Email)
		if exists {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, exists, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalido")
	}

	return claims, false, string(""), err
}

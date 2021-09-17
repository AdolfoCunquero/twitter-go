package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/jwt"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	var usr models.User
	var customErr models.Error

	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {

		http.Error(rw, "Usuario y/o contrasenia invalido", 403)
		return
	}

	if len(usr.Email) == 0 {
		http.Error(rw, "El email del usuario es requerido", 400)
		return
	}

	document, exists := db.Login(usr.Email, usr.Password)

	if !exists {
		customErr.Code = 403
		customErr.Message = "Usuario y/o contrasenia incorrecto"
		utils.JSONResponse(rw, customErr, 403)
		return
	}

	jwtKey, err := jwt.TokenGenerator(document)

	if err != nil {
		http.Error(rw, "Ocurrio un error al generar el token corespondiente "+err.Error(), 403)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	utils.JSONResponse(rw, response, http.StatusCreated)

	// expirationTime := time.Now().Add( 24 * time.Hour)
	// http.SetCookie(rw, &http.Cookie{
	// 	Name:"token",
	// 	Value: jwtKey,
	// 	Expires: expirationTime,
	// })

}

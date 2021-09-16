package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
)

func Register(rw http.ResponseWriter, r *http.Request) {
	var usr models.User
	var customErr models.Error

	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Error en registro de usuario " + err.Error()
		utils.JSONError(rw, customErr, 400)
		return
	}

	if len(usr.Email) == 0 {
		customErr.Code = 400
		customErr.Message = "El Email de usuario es requerido"
		utils.JSONError(rw, customErr, 400)
		return
	}

	if len(usr.Password) < 6 {
		customErr.Code = 400
		customErr.Message = "El Password tiene que tener minimo 6 caracteres"
		utils.JSONError(rw, customErr, 400)
		return
	}
	_, exists, _ := db.UserExists(usr.Email)

	if exists {
		customErr.Code = 400
		customErr.Message = "El email ya existe un usuario registrado con el email " + usr.Email
		utils.JSONError(rw, customErr, 400)
		return
	}

	_, status, err := db.InsertRegister(usr)

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Ocurrio un error al registrar el usuario " + err.Error()
		utils.JSONError(rw, customErr, 400)
		return
	}

	if !status {
		customErr.Code = 400
		customErr.Message = "No se ha logrado registrar el usuario"
		utils.JSONError(rw, customErr, 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

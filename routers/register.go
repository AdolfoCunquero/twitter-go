package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
)

func Register(rw http.ResponseWriter, r *http.Request) {
	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		http.Error(rw, "Error en registro de usuario "+err.Error(), 400)
		return
	}

	if len(usr.Email) == 0 {
		http.Error(rw, "El Email de usuario es requerido", 400)
		return
	}

	if len(usr.Password) < 6 {
		http.Error(rw, "El Password tiene que tener minimo 6 caracteres", 400)
		return
	}
	_, exists, _ := db.UserExists(usr.Email)

	if exists {
		http.Error(rw, "El email ya existe un usuario registrado con el email "+usr.Email, 400)
		return
	}

	_, status, err := db.InsertRegister(usr)

	if err != nil {
		http.Error(rw, "Ocurrio un error al registrar el usuario "+err.Error(), 400)
		return
	}

	if status {
		http.Error(rw, "No se ha logrado registrar el usuario", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

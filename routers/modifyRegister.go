package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
)

func ModifyProfile(rw http.ResponseWriter, r *http.Request) {
	var usr models.User
	var customErr models.Error
	var status bool

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		customErr.Code = 400
		customErr.Message = "Error al decodificar json " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	status, err = db.ModifyRegister(usr, IDUsuario)

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Error al modificar el registro " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	if !status {
		customErr.Code = 400
		customErr.Message = "No se ha logrado modificar el registro del usuario"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	rw.WriteHeader(200)
}

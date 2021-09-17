package routers

import (
	"net/http"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
)

func SearchProfile(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var customErr models.Error

	if len(ID) < 1 {
		customErr.Code = 404
		customErr.Message = "Debe enviar el parametro ID"
		utils.JSONResponse(rw, customErr, 404)
		return
	}

	profile, err := db.SearchProfile(ID)

	if err != nil {
		customErr.Code = 500
		customErr.Message = "Ocurrio un error al buscar el perfil " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	utils.JSONResponse(rw, profile, http.StatusCreated)
}

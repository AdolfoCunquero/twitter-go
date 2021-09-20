package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
	"github.com/gorilla/mux"
)

func DownloadAvatar(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	var customErr models.Error

	if len(userId) < 1 {
		customErr.Code = 400
		customErr.Message = "Debe enviar el parametro ID"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	profile, err := db.SearchProfile(userId)

	if err != nil {
		customErr.Code = 404
		customErr.Message = "Usuario no encontrado"
		utils.JSONResponse(rw, customErr, 404)
		return
	}

	file, err1 := os.Open("uploads/avatar/" + profile.Avatar)
	if err1 != nil {
		customErr.Code = 404
		customErr.Message = "Error al obtener la imagen"
		utils.JSONResponse(rw, customErr, 404)
		return
	}

	_, err = io.Copy(rw, file)
}

func DownloaBanner(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["id"]
	var customErr models.Error

	if len(userId) < 1 {
		customErr.Code = 400
		customErr.Message = "Debe enviar el parametro ID"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	profile, err := db.SearchProfile(userId)

	if err != nil {
		customErr.Code = 404
		customErr.Message = "Usuario no encontrado"
		utils.JSONResponse(rw, customErr, 404)
		return
	}

	file, err1 := os.Open("uploads/banner/" + profile.Banner)
	if err1 != nil {
		customErr.Code = 404
		customErr.Message = "Error al obtener la imagen"
		utils.JSONResponse(rw, customErr, 404)
		return
	}

	_, err = io.Copy(rw, file)

	if err != nil {
		http.Error(rw, "Imagen no encontrada", 404)
		return
	}
}

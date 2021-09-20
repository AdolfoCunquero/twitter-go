package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
)

func UploadAvatar(rw http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var customErr models.Error
	var extension = strings.Split(handler.Filename, ".")[1]
	var image string = "uploads/avatar/" + IDUsuario + "." + extension

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Debe adjuntar un archivo valido " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	f, err1 := os.OpenFile(image, os.O_WRONLY|os.O_CREATE, 0666)

	if err1 != nil {
		customErr.Code = 400
		customErr.Message = "Error al subir imagen " + err1.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	_, err2 := io.Copy(f, file)

	if err2 != nil {
		customErr.Code = 400
		customErr.Message = "Error al subir imagen " + err2.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	var usr models.User
	var status bool

	usr.Avatar = IDUsuario + "." + extension
	status, err = db.ModifyRegister(usr, IDUsuario)

	if !status || err != nil {
		customErr.Code = 400
		customErr.Message = "Error al actualizar datos" + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	rw.WriteHeader(200)
}

func UploadBanner(rw http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("banner")
	var customErr models.Error
	var extension = strings.Split(handler.Filename, ".")[1]
	var image string = "uploads/banner/" + IDUsuario + "." + extension

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Debe adjuntar un archivo valido " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	f, err1 := os.OpenFile(image, os.O_WRONLY|os.O_CREATE, 0666)

	if err1 != nil {
		customErr.Code = 400
		customErr.Message = "Error al subir imagen " + err1.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	_, err2 := io.Copy(f, file)

	if err2 != nil {
		customErr.Code = 400
		customErr.Message = "Error al subir imagen " + err2.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	var usr models.User
	var status bool

	usr.Banner = IDUsuario + "." + extension
	status, err = db.ModifyRegister(usr, IDUsuario)

	if !status || err != nil {
		customErr.Code = 400
		customErr.Message = "Error al actualizar datos " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	rw.WriteHeader(200)
}

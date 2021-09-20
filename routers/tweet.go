package routers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
	"github.com/gorilla/mux"
)

func InsertTweet(rw http.ResponseWriter, r *http.Request) {
	var tweet models.Tweet
	var customErr models.Error
	var err error
	var status bool

	err = json.NewDecoder(r.Body).Decode(&tweet)

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Los datos enviados no son correctos " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	tweet.UserId = IDUsuario
	tweet.Date = time.Now()

	_, status, err = db.InsertTweet(tweet)

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Ocurrio un error al insertar el tweet " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	if !status {
		customErr.Code = 400
		customErr.Message = "No fue posible insertar el tweet"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	rw.WriteHeader(200)
}

func ReadTweets(rw http.ResponseWriter, r *http.Request) {

	var customErr models.Error
	var pageInt int64
	var err error

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		customErr.Code = 400
		customErr.Message = "El parametro id es necesario"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	page := r.URL.Query().Get("page")

	if len(page) < 1 {
		pageInt = 1
	} else {
		pageInt, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			customErr.Code = 400
			customErr.Message = "La pagina debe de ser un numero entero"
			utils.JSONResponse(rw, customErr, 400)
			return
		}
	}

	results, status := db.ReadTweets(ID, pageInt)

	if !status {
		customErr.Code = 400
		customErr.Message = "Error al leer tweets"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	utils.JSONResponse(rw, results, 200)
}

func DeleteTweet(rw http.ResponseWriter, r *http.Request) {
	//tweetId := r.URL.Query().Get("id")

	params := mux.Vars(r)
	tweetId := params["id"]

	var customErr models.Error

	if len(tweetId) < 1 {
		customErr.Code = 400
		customErr.Message = "Debe enviar el parametro id"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	err := db.DeleteTweet(tweetId, IDUsuario)

	if err != nil {
		customErr.Code = 400
		customErr.Message = "Error al eliminar tweet " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	rw.WriteHeader(200)
}

func ReadTweetsFollow(rw http.ResponseWriter, r *http.Request) {
	var customErr models.Error
	var pageInt int64
	var err error

	page := r.URL.Query().Get("page")

	if len(page) < 1 {
		pageInt = 1
	} else {
		pageInt, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			customErr.Code = 400
			customErr.Message = "La pagina debe de ser un numero entero"
			utils.JSONResponse(rw, customErr, 400)
			return
		}
	}

	response, status := db.ReadTweetsFollow(IDUsuario, pageInt)

	if !status {
		customErr.Code = 400
		customErr.Message = "Error al leer los tweets"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	utils.JSONResponse(rw, response, 200)
}

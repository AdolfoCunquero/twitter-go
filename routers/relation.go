package routers

import (
	"net/http"
	"strconv"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/utils"
	"github.com/gorilla/mux"
)

func InsertRelation(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var customErr models.Error

	if len(id) < 1 {
		customErr.Code = 400
		customErr.Message = "El parametro id es obligatorio"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	var relation models.Relation
	relation.UserId = IDUsuario
	relation.UserRelationId = id

	status, err := db.InsertRelation(relation)

	if !status || err != nil {
		customErr.Code = 400
		customErr.Message = "Ocurrio un error al realizar la operacion " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	rw.WriteHeader(200)
}

func DeleteRelation(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var customErr models.Error

	if len(id) < 1 {
		customErr.Code = 400
		customErr.Message = "El parametro id es obligatorio"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	var relation models.Relation
	relation.UserId = IDUsuario
	relation.UserRelationId = id

	status, err := db.DeleteRelation(relation)

	if !status || err != nil {
		customErr.Code = 400
		customErr.Message = "Ocurrio un error al realizar la operacion " + err.Error()
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	rw.WriteHeader(200)
}

func SearchRelation(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	var customErr models.Error
	var response models.RelationExists

	if len(id) < 1 {
		customErr.Code = 400
		customErr.Message = "El parametro id es obligatorio"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	var relation models.Relation
	relation.UserId = IDUsuario
	relation.UserRelationId = id

	status, err := db.SearchRelation(relation)

	if !status || err != nil {
		response.Status = false
	} else {
		response.Status = true
	}

	utils.JSONResponse(rw, response, 200)
}

func ReadAllUsers(rw http.ResponseWriter, r *http.Request) {

	var pageInt int64
	var err error
	var customErr models.Error

	typeUser := r.URL.Query().Get("tipo")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

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

	results, status := db.ReadAllUsers(IDUsuario, pageInt, search, typeUser)

	if !status {
		customErr.Code = 400
		customErr.Message = "Error al leer los usuarios"
		utils.JSONResponse(rw, customErr, 400)
		return
	}

	utils.JSONResponse(rw, results, 200)
}

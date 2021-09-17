package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/AdolfoCunquero/twitter-go/middlew"
	"github.com/AdolfoCunquero/twitter-go/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	//router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.Profile))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3500"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

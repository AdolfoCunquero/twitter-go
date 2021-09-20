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

	router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.SearchProfile))).Methods("GET")
	router.HandleFunc("/profile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")

	router.HandleFunc("/avatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/banner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/avatar/{id}", middlew.CheckDB(middlew.ValidateJWT(routers.DownloadAvatar))).Methods("GET")
	router.HandleFunc("/banner/{id}", middlew.CheckDB(middlew.ValidateJWT(routers.DownloaBanner))).Methods("GET")

	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.InsertTweet))).Methods("POST")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/tweet/{id}", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/tweetHome", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweetsFollow))).Methods("GET")

	router.HandleFunc("/relation/{id}", middlew.CheckDB(middlew.ValidateJWT(routers.InsertRelation))).Methods("POST")
	router.HandleFunc("/relation/{id}", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/relation/{id}", middlew.CheckDB(middlew.ValidateJWT(routers.SearchRelation))).Methods("GET")

	router.HandleFunc("/listUsers", middlew.CheckDB(middlew.ValidateJWT(routers.ReadAllUsers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "3500"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

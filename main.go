package main

import (
	"log"

	"github.com/AdolfoCunquero/twitter-go/db"
	"github.com/AdolfoCunquero/twitter-go/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Sin coenxion a la base de datos")
		return
	}
	handlers.Handlers()
	log.Println("Escuchando en el puerto 3500")
}

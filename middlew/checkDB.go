package middlew

import (
	"net/http"

	"github.com/AdolfoCunquero/twitter-go/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(rw, "Conexion perdida con la DB", 500)
			return
		}
		next.ServeHTTP(rw, r)
	}
}

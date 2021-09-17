package middlew

import (
	"net/http"

	"github.com/AdolfoCunquero/twitter-go/models"
	"github.com/AdolfoCunquero/twitter-go/routers"
	"github.com/AdolfoCunquero/twitter-go/utils"
)

func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ValidateJWT(r.Header.Get("Authorization"))

		if err != nil {
			customErr := models.Error{
				Code:    http.StatusBadRequest,
				Message: "Token invalido " + err.Error(),
			}
			utils.JSONResponse(rw, customErr, http.StatusBadRequest)
			return
		}

		next.ServeHTTP(rw, r)
	}
}

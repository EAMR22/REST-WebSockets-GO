// Este handler se va a encargar de procesar la peticion de la ruta principal.
package handlers

import (
	"encoding/json"
	"net/http"

	"platzi.com/go/rest-ws/server"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // Le dice al cliente que lo que le esta respondiendo es un json.
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Platzi Go",
			Status:  true,
		})
	}
}

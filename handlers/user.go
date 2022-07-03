package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"platzi.com/go/rest-ws/models"
	"platzi.com/go/rest-ws/repository"
	"platzi.com/go/rest-ws/server"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func SignUpHandler(s server.Server) http.HandlerFunc { // El Request es la data que le envia el cliente, y el ResponseWriter es para responderle.
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request) // El Body es el cuerpo de la peticion.
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest) // Especificamos que el error ocurre del lado del cliente.
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError) // Aqui el error ocurre del lado del servidor.
			return
		}

		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			Id:    user.Id,    // Del lado izquierdo el Id es del struct SignUpResponse.
			Email: user.Email, // Del lado derecho el Email lo trae de models.
		})
	}
}

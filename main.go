package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"platzi.com/go/rest-ws/handlers"
	"platzi.com/go/rest-ws/server"
)

func main() {
	err := godotenv.Load(".env") // godotenv es una de las dependencias que le instalamos.

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT") // Las variables tienen que ser iguales a las del archivo .env
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		JWTSecret:   JWT_SECRET, // A la izquierda van los valores del struct de server.go
		Port:        PORT,       // A la derecha los valores del archivo .env
		DatabaseUrl: DATABASE_URL,
	})

	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet) // Es el primer endpoint.
}

package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string // El puerto donde se va a ejecutar.
	JWTSecret   string // La clave secreta para generar tokens.
	DatabaseUrl string // Conexion a la base de datos.
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router // Es un ruteador que va a definir las rutas de la API.
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) { // Es el constructor de la clase
	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}

	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	broker := &Broker{ // Creamos la instancia y Hacemos el retorno del broker.
		config: config,
		router: mux.NewRouter(),
	}

	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) { // Aqui levanta el servidor
	b.router = mux.NewRouter()
	binder(b, b.router) // El broker se comporta como si fuera un servidor.
	log.Println("Starting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

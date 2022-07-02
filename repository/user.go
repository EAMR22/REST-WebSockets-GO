package repository

import (
	"context"

	"platzi.com/go/rest-ws/models"
)

// Patron repository

type UserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id int64) (*models.User, error) // Retorna un usuario o un error en caso de que haya.
	Close() error                                                    // Va a cerrar conexiones a la base de datos.
}

var implementation UserRepository

func setRepository(repository UserRepository) { // UserRepository puede ser una implementacion en mongo db, postgres, etc..
	implementation = repository
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUserById(ctx, id)
}

func Close() error {
	return implementation.Close()
}

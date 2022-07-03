package database

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Con el "_" importa la libreria aunque no la use el programa.
	"platzi.com/go/rest-ws/models"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) { // Es el constructor de la clase.
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repo *PostgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := repo.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)",
		user.Email, user.Password)
	// ExecContext ejecuta una oracion de sql en base de datos con un contexto.
	return err
}

func (repo *PostgresRepository) GetUserById(ctx context.Context, id string) (*models.User, error) {
	rows, err := repo.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	// QueryContext lo que hace es hacer una consulta a la base de datos de postgres.

	defer func() {
		err = rows.Close() // Cerramos el lector para que la base de datos no quede con una conexion abierta.
		if err != nil {
			log.Fatal(nil)
		}
	}()

	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err == nil {
			return &user, nil
		} // Scan lo que hace es copiar las columnas que esta leyendo en el momento, dentro de un destinatario.
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *PostgresRepository) Close() error {
	return repo.db.Close()
}

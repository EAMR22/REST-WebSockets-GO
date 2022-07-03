package models

type User struct {
	Id       string `json:"id"` // En postman hace el cambio a minuscula.
	Email    string `json:"email"`
	Password string `json:"password"`
}

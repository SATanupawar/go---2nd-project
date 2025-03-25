package models 

import (
	"go-postgres-app/config"

)

// user struct represents the user table in the database
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// createUser inserts a new user into the database
func CreateUser(name, email string) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2)"
	_, err := config.DB.Exec(query, name, email)
	return err
}

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

// GetAllUsers retrieves all users from the database
func GetAllUser() ([]User, error) {
	query := "SELECT id , name , email , created_at FROM users ORDER BY id"
	row, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var users []User
	for row.Next() {
		var user User
		err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserById(id string) (User, error) {
	query := "SELECT id , name , email , created_at FROM users WHERE id = $1"
	row, err := config.DB.Query(query, id)
	if err != nil {
		return User{}, err
	}
	defer row.Close()
 
	var user User
	row.Next()
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(id string, name, email string) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := config.DB.Exec(query, name, email, id)
	return err
}

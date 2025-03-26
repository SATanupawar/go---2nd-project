package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go-postgres-app/models"
)

//RequestBody represents the expected JSON structure 
type RequestBody struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUserHandler handles the creation of a new user

func CreateUserhandler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	// Decode the request body into the requestBody struct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the user in the database
	err = models.CreateUser(requestBody.Name, requestBody.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return a success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

// GetAllUsersHandler retrieves all users from the database
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Get all users from the database
	users , err := models.GetAllUser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

// set heder 
w.Header().Set("content-type" , "application/json")

// encode the data 
json.NewEncoder(w).Encode(users)
}




func GetUserByIdHandler(w http.ResponseWriter, r *http.Request){

	id := r.URL.Path[len("/api/users/"):]

	user , err := models.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// set header 
	w.Header().Set("content-type" , "application/json")

	// encode the data 
	json.NewEncoder(w).Encode(user)

}
                                                              
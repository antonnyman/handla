package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// User struct
type User struct {
	gorm.Model
	Firstname string
	Lastname  string
	Email     string
	Items     []Item
}

// UserRoutes returns all user routes
func UserRoutes(r *mux.Router) {
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
}

// UserMigration creates the table or reads from existing
func UserMigration() {
	db.AutoMigrate(&User{})
}

// GetUsers returns all users from the db
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// CreateUser creates a user from a post request
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)

	if db.Where("Email = ?", user.Email).First(&user).RecordNotFound() {
		db.Create(&User{
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
			Email:     user.Email,
		})

		json.NewEncoder(w).Encode(db.Last(&user).Value)
		return
	}

	fmt.Fprintf(w, `{"error": "User with email %s already exists."}`, user.Email)
}

// GetUser returns a single user
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	db.Where("ID = ?", id).Find(&user)

	json.NewEncoder(w).Encode(user)
}

// UpdateUser updates a single user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	db.Where("ID = ?", id).Find(&user)

	var updateUser User
	_ = json.NewDecoder(r.Body).Decode(&updateUser)

	user.Email = updateUser.Email
	user.Firstname = updateUser.Firstname
	user.Lastname = updateUser.Lastname

	db.Save(&user)

	json.NewEncoder(w).Encode(user)
}

// DeleteUser pernamently deletes a user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	db.Where("ID = ?", id).Find(&user)
	db.Delete(&user)
	fmt.Fprintf(w, `{"notice": "User %s was deleted."}`, user.Email)
}

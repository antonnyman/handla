package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type List struct {
	List   string
	ListID string
}

// ListRoutes returns all routes for lists
func ListRoutes(r *mux.Router) {
	r.HandleFunc("/api/lists", GetAllLists).Methods("GET")
	r.HandleFunc("/api/lists/{id}", GetListByID).Methods("GET")
}

// GetAllLists returns list names and ids
func GetAllLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var lists []List
	var items []Item

	db.Find(&items)

	for _, item := range items {
		lists = append(lists, List{List: item.List, ListID: item.ListID})
	}
	json.NewEncoder(w).Encode(lists)
}

// GetListByID returns all lists in an array
func GetListByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items []Item

	vars := mux.Vars(r)
	id := vars["id"]

	db.Where("list_id = ?", id).Find(&items)
	json.NewEncoder(w).Encode(items)
}

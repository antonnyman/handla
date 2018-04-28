package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Item struct
type Item struct {
	gorm.Model
	Name       string
	Count      int
	Store      string
	List       string
	Checked    bool
	AssignedTo int
	UserID     int
}

// ItemRoutes returns all item routes
func ItemRoutes(r *mux.Router) {
	r.HandleFunc("/api/items", GetItems).Methods("GET")
	r.HandleFunc("/api/items/checked", GetCheckedItems).Methods("GET")
	r.HandleFunc("/api/items/unchecked", GetUncheckedItems).Methods("GET")
	r.HandleFunc("/api/user/{id}/items", GetItemsByUserID).Methods("GET")
	r.HandleFunc("/api/items/{id}", GetItem).Methods("GET")
	r.HandleFunc("/api/items", CreateItem).Methods("POST")
	r.HandleFunc("/api/items/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}/check", CheckItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}/uncheck", CheckItem).Methods("PUT")
	r.HandleFunc("/api/items/{id}", DeleteItem).Methods("DELETE")
}

// ItemMigration creates the table or reads from existing
func ItemMigration() {
	db.AutoMigrate(&Item{})
}

// GetItems returns all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items []Item
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}

// GetCheckedItems returns all checked items
func GetCheckedItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items []Item
	db.Where("Checked = ?", true).Find(&items)
	json.NewEncoder(w).Encode(items)
}

// GetUncheckedItems returns all unchecked items
func GetUncheckedItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items []Item
	db.Where("Checked = ?", false).Find(&items)
	json.NewEncoder(w).Encode(items)
}

// GetItemsByUserID returns all items
func GetItemsByUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var items []Item
	var user User

	vars := mux.Vars(r)
	id := vars["id"]

	db.Where("ID = ?", id).Find(&user)
	db.Model(&user).Related(&items)

	json.NewEncoder(w).Encode(items)
}

// GetItem returns a specific item
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var item Item
	db.Where("ID = ?", id).Find(&item)
	json.NewEncoder(w).Encode(item)
}

// CreateItem creates an item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	db.Create(&Item{
		Name:       item.Name,
		Count:      item.Count,
		Store:      item.Store,
		List:       item.List,
		Checked:    item.Checked,
		AssignedTo: item.AssignedTo,
		UserID:     item.UserID,
	})

	json.NewEncoder(w).Encode(db.Last(&item).Value)
}

// UpdateItem updates all values on a specific item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var item Item
	db.Where("ID = ?", id).Find(&item)

	var updateItem Item
	_ = json.NewDecoder(r.Body).Decode(&updateItem)

	item.Name = updateItem.Name
	item.Store = updateItem.Store
	item.List = updateItem.List
	item.Checked = updateItem.Checked
	item.AssignedTo = updateItem.AssignedTo
	item.UserID = updateItem.UserID

	db.Save(&item)

	json.NewEncoder(w).Encode(item)
}

// CheckItem checks or unchecks a specific item
func CheckItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var item Item
	db.Where("ID = ?", id).Find(&item)

	item.Checked = !item.Checked

	db.Save(&item)

	json.NewEncoder(w).Encode(item)
}

// DeleteItem deletes a specific item
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]

	var item Item
	db.Where("ID = ?", id).Find(&item)
	db.Delete(&item)
	fmt.Fprintf(w, `{"notice": "Item %s was deleted."}`, item.Name)

}

package main

/*

 === NOT USED ===

// List struct
type List struct {
	gorm.Model
	Name  string
	Items []Item
}

// ListRoutes returns all routes for lists
func ListRoutes(r *mux.Router) {
	r.HandleFunc("/api/lists", GetLists).Methods("GET")
	r.HandleFunc("/api/lists/{id}", GetList).Methods("GET")
}

// ListMigration creates the list table or reads from it if it exists
func ListMigration() {
	db.AutoMigrate(&List{})
}

func CreateListRelations() {
	var list List
	var item Item
	db.Model(&list).Related(&item)
}

// GetLists returns all lists
func GetLists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var lists []List
	db.Find(&lists)
	json.NewEncoder(w).Encode(lists)
}

// GetList returns a specific list
func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	var list List
	db.Where("ID = ?", id).Find(&list)
	json.NewEncoder(w).Encode(list)
}
*/

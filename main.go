package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func initDB() {
	db, err = gorm.Open("sqlite3", "handla.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
}

func main() {
	initDB()
	UserMigration()
	ItemMigration()

	r := mux.NewRouter()
	UserRoutes(r)
	ItemRoutes(r)
	ListRoutes(r)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:*", "https://*.now.sh"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

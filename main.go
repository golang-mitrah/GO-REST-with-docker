package main

import (
	"farm_management/database"
	"farm_management/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Initialize Database
	database.Connect()
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter()
	routes.RegisterProductRoutes(router)

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", 8020), router))
}

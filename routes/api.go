package routes

import (
	"farm_management/controllers"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/products", controllers.CreateProduct).Methods("POST")
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/login", controllers.UserLogin).Methods("POST")
	r.HandleFunc("/products/{id}", controllers.GetProductById).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.UpdateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}

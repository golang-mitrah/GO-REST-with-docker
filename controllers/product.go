package controllers

import (
	"encoding/json"
	"farm_management/database"
	"farm_management/entities"
	"farm_management/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.Users
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Find(&user, "email = ?", user.Email)
	loginResponse, _ := middlewares.CreateToken(user.Email)
	json.NewEncoder(w).Encode(loginResponse)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entities.Users
	json.NewDecoder(r.Body).Decode(&user)
	database.Instance.Table("users").Create(&user)
	json.NewEncoder(w).Encode(user)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	_, err := middlewares.ExtractTokenID(r)
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var product entities.Product
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Create(&product)
	json.NewEncoder(w).Encode(product)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	_, err := middlewares.ExtractTokenID(r)
	if err != nil {
		json.NewEncoder(w).Encode("Unauthorized")
		return
	}
	var products []entities.Product
	database.Instance.Find(&products)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.First(&product, productId)
	json.NewDecoder(r.Body).Decode(&product)
	database.Instance.Save(&product)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productId := mux.Vars(r)["id"]
	if !checkIfProductExists(productId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var product entities.Product
	database.Instance.Delete(&product, productId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func checkIfProductExists(productId string) bool {
	var product entities.Product
	database.Instance.First(&product, productId)
	return product.ID != 0
}

package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Youngprinnce/shop-anything-lagos/models"
	"github.com/Youngprinnce/shop-anything-lagos/services"
	"github.com/Youngprinnce/shop-anything-lagos/utils"
	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchantID"]
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	newProduct, err := services.CreateProduct(merchantID, product)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}
	utils.Respond(w, newProduct, http.StatusCreated)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchantID"]

	products, err := services.GetAllProducts(merchantID)
	if err != nil {
		utils.HandleError(w, err, http.StatusNotFound)
		return
	}
	utils.Respond(w, products, http.StatusOK)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchantID"]
	skuID := mux.Vars(r)["skuID"]

	product, err := services.GetProduct(merchantID, skuID)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}
	utils.Respond(w, product, http.StatusOK)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchantID"]
	skuID := mux.Vars(r)["skuID"]

	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	updatedProduct, err := services.UpdateProduct(merchantID, skuID, product)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}
	utils.Respond(w, updatedProduct, http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchantID"]
	skuID := mux.Vars(r)["skuID"]

	err := services.DeleteProduct(merchantID, skuID)
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}
	utils.Respond(w, nil, http.StatusNoContent)
}

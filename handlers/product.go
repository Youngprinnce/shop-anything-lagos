package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Youngprinnce/shop-anything-lagos/models"
	"github.com/Youngprinnce/shop-anything-lagos/services"
	"github.com/Youngprinnce/shop-anything-lagos/utils"
	"github.com/gorilla/mux"
)

func createProduct(w http.ResponseWriter, r *http.Request) {
	merchantID := mux.Vars(r)["merchant_id"]
	var product models.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		utils.HandleError(w, err)
		return
	}

	newProduct, err := services.CreateProduct(merchantID, product)
	if err != nil {
		utils.HandleError(w, err)
		return
	}
	utils.Respond(w, newProduct, http.StatusCreated)
}

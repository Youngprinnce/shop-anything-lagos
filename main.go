package main

import (
	"log"
	"net/http"

	"github.com/Youngprinnce/shop-anything-lagos/config"
	"github.com/Youngprinnce/shop-anything-lagos/handlers"
	"github.com/Youngprinnce/shop-anything-lagos/middleware"
	"github.com/Youngprinnce/shop-anything-lagos/utils"

	"github.com/gorilla/mux" // Third-party router
)

func main() {
    router := mux.NewRouter()

	router.HandleFunc("/api/products/{merchantID}", middleware.ValidateProduct(handlers.CreateProduct, utils.ValidateCreateProduct)).Methods(http.MethodPost)
	router.HandleFunc("/api/products/{merchantID}", handlers.GetAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{merchantID}/{skuID}", handlers.GetProduct).Methods(http.MethodGet)
	router.HandleFunc("/api/products/{merchantID}/{skuID}", middleware.ValidateProduct(handlers.UpdateProduct, utils.ValidateUpdateProduct)).Methods(http.MethodPut)
	router.HandleFunc("/api/products/{merchantID}/{skuID}", handlers.DeleteProduct).Methods(http.MethodDelete)

    log.Println("Server running on port", config.ServerAddr)
    log.Fatal(http.ListenAndServe(config.ServerAddr, router))
}

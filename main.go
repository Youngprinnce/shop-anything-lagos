package main

import (
	"log"
	"net/http"

	"github.com/Youngprinnce/shop-anything-lagos/config"
	"github.com/gorilla/mux" // Third-party router
)

func main() {
    router := mux.NewRouter()

    log.Println("Server running on port", config.ServerAddr)
    log.Fatal(http.ListenAndServe(config.ServerAddr, router))
}

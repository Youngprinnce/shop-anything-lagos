package utils

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

func HandleError(w http.ResponseWriter, err error, statusCode int) {
    log.Println(err)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    var message string
    if statusCode == http.StatusInternalServerError {
        message = "Internal server error"
    } else {
        message = err.Error()
    }
    json.NewEncoder(w).Encode(map[string]string{"error": message, "status": strconv.Itoa(statusCode)})
}

func NewError(message string) error {
    return errors.New(message)
}
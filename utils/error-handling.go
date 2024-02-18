package utils

import (
    "log"
    "net/http"
)

func HandleError(w http.ResponseWriter, err error) {
    log.Println(err)
    http.Error(w, err.Error(), http.StatusInternalServerError)
}
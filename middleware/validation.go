package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/Youngprinnce/shop-anything-lagos/models"
	"github.com/Youngprinnce/shop-anything-lagos/utils"
)

func ValidateProduct(next http.HandlerFunc, validate func(models.Product) error) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var product models.Product
        err := json.NewDecoder(r.Body).Decode(&product)
        if err != nil {
            utils.HandleError(w, err, http.StatusBadRequest)
            return
        }

        err = validate(product)
        if err != nil {
            utils.HandleError(w, err, http.StatusBadRequest)
            return
        }

        next(w, r)
    }
}

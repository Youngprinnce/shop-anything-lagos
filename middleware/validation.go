package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Youngprinnce/shop-anything-lagos/models"
	"github.com/Youngprinnce/shop-anything-lagos/utils"
)

func ValidateProduct(next http.HandlerFunc, validate func(models.Product) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			utils.HandleError(w, err, http.StatusInternalServerError)
			return
		}

		// Restore the request body with the original data
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Decode the request body into a Product struct for validation
		var product models.Product
		if err := json.Unmarshal(body, &product); err != nil {
			utils.HandleError(w, err, http.StatusBadRequest)
			return
		}

		// Validate the product
		if err := validate(product); err != nil {
			utils.HandleError(w, err, http.StatusBadRequest)
			return
		}

		// Call the next handler
		next(w, r)
	}
}

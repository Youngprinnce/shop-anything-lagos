package utils

import (
	"errors"

	"github.com/Youngprinnce/shop-anything-lagos/models"
)

// creation a new function to validate the product
func ValidateCreateProduct(product models.Product) error {
	if product.SKU == "" {
        return errors.New("SKU is required")
    }
	if product.Name == "" {
		return errors.New("Name is required")
	}
	if product.Description == "" {
		return errors.New("Description is required")
	}
	if product.Price <= 0.0 {
		return errors.New("Price must be positive")
	}
	return nil
}
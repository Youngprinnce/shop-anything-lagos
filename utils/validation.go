package utils

import (
	"errors"

	"github.com/Youngprinnce/shop-anything-lagos/models"
)

// creation a new function to validate the product
func ValidateCreateProduct(product models.Product) error {
	if product.Sku == "" {
        return errors.New("sku is required")
    }
	if product.Name == "" {
		return errors.New("name is required")
	}
	if product.Description == "" {
		return errors.New("description is required")
	}
	if product.Price <= 0.0 {
		return errors.New("price must be positive")
	}
	return nil
}

// creation a new function to validate the product
func ValidateUpdateProduct(product models.Product) error {
	if product.Name == "" {
		return errors.New("name is required")
	}
	if product.Description == "" {
		return errors.New("description is required")
	}
	if product.Price <= 0.0 {
		return errors.New("price must be positive")
	}
	return nil
}
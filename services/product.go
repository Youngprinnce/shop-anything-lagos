package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/Youngprinnce/shop-anything-lagos/models"
	"github.com/Youngprinnce/shop-anything-lagos/utils"
)

var (
	// Use a sync.Map for storing products with merchant ID as key and a slice of products as value
	products = sync.Map{}
)

// CreateProduct creates a new product for a merchant.
func CreateProduct(merchantID string, product models.Product) (models.Product, error) {
	// Get the products for the merchant
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		// If the merchant has no products, create a new slice
		merchantProducts = []models.Product{}
	}

	//  Check if the product sku already exists in the merchant's products
	for _, merchantProduct := range merchantProducts.([]models.Product) {
		if merchantProduct.SKU == product.SKU {
			return models.Product{}, fmt.Errorf("product with SKU %s already exists for merchant %s", product.SKU, merchantID)
		}
	}

	// Generate a new product ID
	productID, err := utils.GenerateProductID()
	if err != nil {
		return models.Product{}, err
	}
	product.ID = productID
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.MerchantID = merchantID

	// Append the new product to the slice
	products.Store(merchantID, append(merchantProducts.([]models.Product), product))

	// console the content of the products to the console
	// products.Range(func(key, value interface{}) bool {
	// 	fmt.Printf("Merchant ID: %v\n", key)
	// 	product := value.([]models.Product)
	// 	for _, p := range product {
	// 		fmt.Printf("Product: %v\n", p)
	// 	}
	// 	return true
	// })
	return product, nil
}

// GetAllProducts gets all products for a merchant.
func GetAllProducts(merchantID string) ([]models.Product, error) {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return []models.Product{}, nil
	}

	return merchantProducts.([]models.Product), nil
}

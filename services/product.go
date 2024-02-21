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
	merchantProducts, _ := products.LoadOrStore(merchantID, make(map[string]models.Product))

	//  Check if the product sku already exists in the merchant's products
	if _, exists := merchantProducts.(map[string]models.Product)[product.Sku]; exists {
		return models.Product{}, fmt.Errorf("product with SKU %s already exists for merchant %s", product.Sku, merchantID)
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

	// Store the product in the merchants product map using sku as key
	merchantProducts.(map[string]models.Product)[product.Sku] = product

	return product, nil
}

// GetAllProducts gets all products for a merchant.
func GetAllProducts(merchantID string) (map[string]models.Product, error) {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return merchantProducts.(map[string]models.Product), nil
	}

	return merchantProducts.(map[string]models.Product), nil
}

// GetProduct gets a product by its SKU and merchant ID.
func GetProduct(merchantID string, skuID string) (models.Product, error) {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return models.Product{}, utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
	}

	// Retrieve the product by SKU from the merchant's product map
	if product, exists := merchantProducts.(map[string]models.Product)[skuID]; exists {
		return product, nil
	}

	return models.Product{}, utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
}

// UpdateProduct updates a product by its SKU and merchant ID.
func UpdateProduct(merchantID string, skuID string, product models.Product) (models.Product, error) {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return models.Product{}, utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
	}

	var existingProduct models.Product

	// Check if the product exist
	existingProduct, exists := merchantProducts.(map[string]models.Product)[skuID];
	
	if !exists {
		return models.Product{}, utils.NewError(fmt.Sprintf("product with SKU %s not found for merchant %s", skuID, merchantID))
	}
	
	// Only update name, decsription and price and update the updated at field
	existingProduct.Name = product.Name
	existingProduct.Description = product.Description
	existingProduct.Price = product.Price
	existingProduct.UpdatedAt = time.Now()

	// Update the product in the slice
	merchantProducts.(map[string]models.Product)[skuID] = existingProduct

	return existingProduct, nil
}

// DeleteProduct deletes a product by its SKU and merchant ID.
func DeleteProduct(merchantID string, skuID string) error {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
	}

	// Check if the product exists
	if _, exists := merchantProducts.(map[string]models.Product)[skuID]; !exists {
		return utils.NewError(fmt.Sprintf("product with SKU %s not found for merchant %s", skuID, merchantID))
	}

	// Delete the product from the merchant's products map
	delete(merchantProducts.(map[string]models.Product), skuID)

	return nil
}
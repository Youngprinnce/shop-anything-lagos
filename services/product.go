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

// GetProduct gets a product by its SKU and merchant ID.
func GetProduct(merchantID string, skuID string) (models.Product, error) {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return models.Product{}, utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
	}

	for _, product := range merchantProducts.([]models.Product) {
		if product.SKU == skuID {
			return product, nil
		}
	}

	return models.Product{}, utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
}

// UpdateProduct updates a product by its SKU and merchant ID.
func UpdateProduct(merchantID string, skuID string, product models.Product) (models.Product, error) {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return models.Product{}, utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
	}

	for i, p := range merchantProducts.([]models.Product) {
		if p.SKU == skuID {
			// Only update name, decsription and price and update the updated at field
			p.Name = product.Name
			p.Description = product.Description
			p.Price = product.Price
			p.UpdatedAt = time.Now()

			// Update the product in the slice
			merchantProducts.([]models.Product)[i] = p
			products.Store(merchantID, merchantProducts)

			return p, nil
		}
	}

	return models.Product{}, utils.NewError(fmt.Sprintf("product with SKU %s not found for merchant %s", skuID, merchantID))
}

// DeleteProduct deletes a product by its SKU and merchant ID.
func DeleteProduct(merchantID string, skuID string) error {
	merchantProducts, ok := products.Load(merchantID)
	if !ok {
		return utils.NewError(fmt.Sprintf("no products found for merchant %s", merchantID))
	}

	for i, p := range merchantProducts.([]models.Product) {
		if p.SKU == skuID {
			// Remove the product from the slice
			// First argument in the append method creates a new slice containing elements from the original slice up to (but not including) index i.
			// Second argument in the append method creates a new slice containing elements from the original slice starting from index i+1 to the end.
			// The ellipsis operator (...) is used to unpack the elements of the two slices.
			newMerchantProducts := append(merchantProducts.([]models.Product)[:i], merchantProducts.([]models.Product)[i+1:]...)
			products.Store(merchantID, newMerchantProducts)

			return nil
		}
	}

	return utils.NewError(fmt.Sprintf("product with SKU %s not found for merchant %s", skuID, merchantID))
}
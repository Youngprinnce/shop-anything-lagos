package services_test

import (
	"testing"
	//"time"

	"github.com/Youngprinnce/shop-anything-lagos/models"
	"github.com/Youngprinnce/shop-anything-lagos/services"
)

func TestCreateProduct(t *testing.T) {
	merchantID := "1"
	product := models.Product{
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       1000,
		Sku:         "SKU-1",
	}

	// Create product
	createdProduct, err := services.CreateProduct(merchantID, product)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
	}

	// Check if product was created successfully
	if createdProduct.ID == "" {
		t.Error("Expected product ID to be set, got empty string")
	}

	// Check if timestamps are set
	if createdProduct.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set, got zero value")
	}
	if createdProduct.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set, got zero value")
	}

	// Check if merchantID is set correctly
	if createdProduct.MerchantID != merchantID {
		t.Errorf("Expected MerchantID to be %s, got %s", merchantID, createdProduct.MerchantID)
	}
}

func TestGetAllProducts(t *testing.T) {
	merchantID := "1"
	product := models.Product{
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       1000,
		Sku:         "SKU-1",
	}

	// Create product
	createdProduct, err := services.CreateProduct(merchantID, product)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
	}

	// Get all products
	products, err := services.GetAllProducts(merchantID)
	if err != nil {
		t.Errorf("Error getting products: %v", err)
	}

	// Check if product was returned
	if len(products) == 0 {
		t.Error("Expected products to be returned, got empty slice")
	}

	// Check if created product was returned
	if products[product.Sku].Sku != createdProduct.Sku {
		t.Errorf("Expected product with Sku %s to be returned, got product with Sku %s", createdProduct.Sku, products[product.Sku].Sku)
	}
}

func TestGetProduct(t *testing.T) {
	merchantID := "1"
	product := models.Product{
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       1000,
		Sku:         "SKU-1",
	}

	// Create product
	createdProduct, err := services.CreateProduct(merchantID, product)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
	}

	// Get product
	getProduct, err := services.GetProduct(merchantID, createdProduct.Sku)
	if err != nil {
		t.Errorf("Error getting product: %v", err)
	}

	// Check if product was returned
	if getProduct.ID != createdProduct.ID {
		t.Errorf("Expected product with ID %s to be returned, got product with ID %s", createdProduct.ID, product.ID)
	}
}

func TestUpdateProduct(t *testing.T) {
	merchantID := "1"
	product := models.Product{
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       1000,
		Sku:         "SKU-1",
	}

	// Create product
	createdProduct, err := services.CreateProduct(merchantID, product)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
	}

	// Update product
	product.Name = "Product 1 updated"
	product.Description = "Product 1 description updated"
	product.Price = 20.50
	updatedProduct, err := services.UpdateProduct(merchantID, createdProduct.Sku, product)
	if err != nil {
		t.Errorf("Error updating product: %v", err)
	}

	// Check if product was updated
	if updatedProduct.Name != product.Name {
		t.Errorf("Expected product name to be %s, got %s", product.Name, updatedProduct.Name)
	}
	if updatedProduct.Description != product.Description {
		t.Errorf("Expected product description to be %s, got %s", product.Description, updatedProduct.Description)
	}
}

func TestDeleteProduct(t *testing.T) {
	merchantID := "1"
	product := models.Product{
		Name:        "Product 1",
		Description: "Product 1 description",
		Price:       1000,
		Sku:         "SKU-1",
	}

	// Create product
	createdProduct, err := services.CreateProduct(merchantID, product)
	if err != nil {
		t.Errorf("Error creating product: %v", err)
	}

	// Delete product
	err = services.DeleteProduct(merchantID, createdProduct.Sku)
	if err != nil {
		t.Errorf("Error deleting product: %v", err)
	}

	// Get product
	_, err = services.GetProduct(merchantID, createdProduct.Sku)
	if err == nil {
		t.Error("Expected error getting product, got nil")
	}
}
package models

import (
	"time"
)

// Product represents a product in the shop
type Product struct {
    // Define product fields here
    ID          string `json:"id,omitempty" gorm:"primary_key"`
    MerchantID  string `json:"merchant_id" gorm:"index"`
    Sku         string `json:"sku" gorm:"unique"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Price       float64 `json:"price"`
    CreatedAt   time.Time `json:"created_at,omitempty" gorm:"created_at"`
    UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"updated_at"`
}

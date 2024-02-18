package utils

import (
    "errors"
    "math/rand"
    "strings"
    "time"
)

var (
    // Global map for storing generated IDs
    generatedIDs map[string]bool = make(map[string]bool)
)

func init() {
    rand.Seed(time.Now().UnixNano()) // Seed random number generator
}

// GenerateProductID generates a unique 16-character alphanumeric product ID.
func GenerateProductID() (string, error) {
    // Base characters for ID generation
    const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

    for {
        // Generate ID with random characters
        var b strings.Builder
        for i := 0; i < 16; i++ {
            b.WriteByte(charset[rand.Intn(len(charset))])
        }
        productID := b.String()

        // Check for uniqueness using the global map
        if !generatedIDs[productID] {
            generatedIDs[productID] = true
            return productID, nil
        }
    }

    // Unreachable in practice due to the vast number of possible IDs
    return "", errors.New("failed to generate unique product ID")
}

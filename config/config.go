package config

import (
    "fmt"
    "os"
    "strconv"
)

// Server configuration variables
var (
    ServerAddr string
    ServerPort int
)

func init() {
    // Load configuration from environment variables
    ServerAddr = getEnvString("SERVER_ADDR", "localhost:8080")
	_, err := getEnvInt("SERVER_PORT", 8080)
    if err != nil {
        fmt.Println("Error reading server port:", err)
        os.Exit(1)
    }
}

func getEnvString(key string, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

func getEnvInt(key string, defaultValue int) (int, error) {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue, nil
    }
    return strconv.Atoi(value)
}

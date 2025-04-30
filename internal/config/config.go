package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var apiKey string

func SetAPIKey(key string) {
	apiKey = key
}

func APIKey() string {
	return apiKey
}

func LoadConfig() error {

	if envKey := os.Getenv("NASA_API_KEY"); envKey != "" {
		apiKey = envKey
		fmt.Println("loaded API key from environment variable")
		return nil
	}

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file path")
	}

	configPath := filepath.Join(filepath.Dir(currentFile), "..", "..", "internal", "config", "config.json")
	fmt.Println("looking for config file at:", configPath)

	file, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close config file: %v\n", err)
		}
	}()

	var config struct {
		APIKey string `json:"api_key"`
	}

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return fmt.Errorf("failed to decode config file: %w", err)
	}

	if config.APIKey == "" {
		return fmt.Errorf("api_key is empty in config file")
	}

	apiKey = config.APIKey
	fmt.Println("loaded API key from config file")
	return nil
}

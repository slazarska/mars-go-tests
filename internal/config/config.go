package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/slazarska/mars-go-tests/internal/log"
)

var apiKey string

func APIKey() string {
	return apiKey
}

func SetAPIKey(key string) {
	apiKey = key
}

func SetTestAPIKey() {
	SetAPIKey("test_key")
}

func SetupRealAPIKey() error {
	if err := LoadConfig(); err != nil {
		log.Error("failed to load config", "error", err)
		return err
	}
	log.Info("API key was set successfully")
	return nil
}

func LoadConfig() error {
	if envKey := os.Getenv("NASA_API_KEY"); envKey != "" {
		apiKey = envKey
		log.Info("loaded API key from environment variable")
		return nil
	}

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file path")
	}

	configPath := filepath.Join(filepath.Dir(currentFile), "..", "..", "internal", "config", "config.json")
	log.Info("looking for config file at:", "path", configPath)

	file, err := os.Open(configPath)
	if err != nil {
		log.Error("failed to open config file", "error", err)
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var config struct {
		APIKey string `json:"api_key"`
	}

	if err := json.NewDecoder(file).Decode(&config); err != nil {
		log.Error("failed to decode config file", "error", err)
		return fmt.Errorf("failed to decode config file: %w", err)
	}

	if config.APIKey == "" {
		log.Error("api_key is empty in config file")
		return fmt.Errorf("api_key is empty in config file")
	}

	apiKey = config.APIKey
	log.Info("loaded API key from config file")
	return nil
}

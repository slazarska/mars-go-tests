package api

import (
	"encoding/json"
	"fmt"
	"github.com/slazarska/mars-go-tests/internal/config"
	"github.com/slazarska/mars-go-tests/internal/constants"
	"github.com/slazarska/mars-go-tests/internal/models"
	"io"
	"log"
	"net/http"
)

// URL для запросов (по умолчанию будет использоваться константа)
var BaseURL = constants.BaseURL

func SetAPIKey(key string) {
	config.SetAPIKey(key)
}

// Функция для получения фотографий с Марса, с возможностью передать свой URL для тестов
func GetMarsPhotos(rover, camera, solValue string, customURL ...string) (*models.PhotoResponse, error) {
	apiKey := config.APIKey()
	if apiKey == "" {
		return nil, fmt.Errorf("API key is missing")
	}

	// Если передан customURL, используем его, иначе используем BaseURL
	url := fmt.Sprintf(BaseURL, rover, solValue, camera, apiKey)
	if len(customURL) > 0 {
		url = fmt.Sprintf(customURL[0], rover, solValue, camera, apiKey)
	}

	log.Printf("request URL: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mars photos: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %v\n", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("unexpected status: %d, body: %s", resp.StatusCode, string(body))
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result models.PhotoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &result, nil
}

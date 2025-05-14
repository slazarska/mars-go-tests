package api

import (
	"encoding/json"
	"fmt"
	"github.com/slazarska/mars-go-tests/internal/config"
	"github.com/slazarska/mars-go-tests/internal/constants"
	"github.com/slazarska/mars-go-tests/internal/log"
	"github.com/slazarska/mars-go-tests/internal/models"
	"io"
	"net/http"
)

func GetMarsPhotos(rover, camera, solValue string, customURL ...string) (*models.RoverResponse, error) {
	apiKey := config.APIKey()
	if apiKey == "" {
		return nil, fmt.Errorf("API key is missing")
	}

	url := fmt.Sprintf(constants.BaseURL, rover, solValue, camera, apiKey)
	if len(customURL) > 0 {
		url = fmt.Sprintf(customURL[0], rover, solValue, camera, apiKey)
	}

	log.Info("sending request", "url", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mars photos: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error("failed to close response body", "error", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		body := string(bodyBytes)
		log.Error("unexpected response",
			"status", resp.StatusCode,
			"body", body,
		)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	var result models.RoverResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &result, nil
}

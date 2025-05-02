package test

import (
	"encoding/json"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/config"
	"github.com/slazarska/mars-go-tests/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetTestAPIKey() {
	config.SetAPIKey("test_key")
}

func TestMockGetMarsPhotos(t *testing.T) {
	SetTestAPIKey()

	// Создаем мок-ответ
	mockResponse := models.PhotoResponse{
		Photos: []models.Photo{
			{
				ID:        12345,
				Sol:       1000,
				ImgSrc:    "http://example.com/image.jpg",
				EarthDate: "2015-06-03",
				Camera: models.Camera{
					Name:     "FHAZ",
					FullName: "Front Hazard Avoidance Camera",
				},
				Rover: models.Rover{
					Name: "Curiosity",
				},
			},
		},
	}

	// Мок-сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		err := json.NewEncoder(w).Encode(mockResponse)
		if err != nil {
			return
		}
	}))
	defer server.Close()

	// Локально определяем BaseURL для теста
	mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"

	// Вызываем API метод с этим мок-URL
	result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result.Photos))
	assert.Equal(t, 12345, result.Photos[0].ID)
}

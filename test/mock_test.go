package test

import (
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/models"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockGetMarsPhotos(t *testing.T) {
	SetTestAPIKey()

	t.Run("success with one photo", func(t *testing.T) {
		var mockResponse models.RoverResponse
		LoadMockJSON(t, "mock_response.json", &mockResponse)

		server := createMockServer(t, http.StatusOK, mockResponse, true)
		defer server.Close()

		mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"

		result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, 1, len(result.Photos))
		assert.Equal(t, 12345, result.Photos[0].ID)
		assert.Equal(t, "http://example.com/image.jpg", result.Photos[0].ImgSrc)
		assert.Equal(t, "Curiosity", result.Photos[0].Rover.Name)
	})

	t.Run("success with empty photo list", func(t *testing.T) {
		mockResponse := models.RoverResponse{
			Photos: []models.Photo{},
		}

		server := createMockServer(t, http.StatusOK, mockResponse, true)
		defer server.Close()

		mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"

		result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Empty(t, result.Photos)
	})

	t.Run("internal server error", func(t *testing.T) {
		server := createMockServer(t, http.StatusInternalServerError, nil, false)
		defer server.Close()

		mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"

		result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

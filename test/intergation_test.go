package test

import (
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupRealAPIKey(t *testing.T) {
	t.Helper()

	err := config.LoadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}
	api.SetAPIKey(config.APIKey())
}

func TestGetMarsPhotos(t *testing.T) {
	setupRealAPIKey(t)

	result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, len(result.Photos), 0, "expected photos, but got none")
}

func TestGetMarsPhotosByRover(t *testing.T) {
	setupRealAPIKey(t)

	rovers := []string{"curiosity", "opportunity", "spirit"}

	for _, rover := range rovers {
		rover := rover
		t.Run("Rover: "+rover, func(t *testing.T) {
			result, err := api.GetMarsPhotos(rover, "fhaz", "100")

			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Greater(t, len(result.Photos), 0, "no photos returned for rover "+rover)
		})
	}
}

func TestGetMarsPhotosByCamera(t *testing.T) {
	setupRealAPIKey(t)

	cameras := []string{"fhaz", "rhaz", "mast"}

	for _, camera := range cameras {
		camera := camera
		t.Run("Camera: "+camera, func(t *testing.T) {
			result, err := api.GetMarsPhotos("curiosity", camera, "1000")

			assert.NoError(t, err)
			assert.NotNil(t, result)
			assert.Greater(t, len(result.Photos), 0, "no photos returned for camera "+camera)
		})
	}
}

func TestGetMarsPhotos_InvalidRover_ReturnsError(t *testing.T) {
	setupRealAPIKey(t)

	result, err := api.GetMarsPhotos("NonExistingRover", "fhaz", "1000")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetMarsPhotos_InvalidCamera_ReturnsEmptyList(t *testing.T) {
	setupRealAPIKey(t)

	result, err := api.GetMarsPhotos("opportunity", "NonExistingCamera", "1000")

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result.Photos, "expected an empty list, but got photos")
}

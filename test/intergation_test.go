package test

import (
	"github.com/slazarska/mars-go-tests/internal/api"
	data "github.com/slazarska/mars-go-tests/internal/constants"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMarsPhotosAllRoversAllCameras(t *testing.T) {
	SetupRealAPIKey(t)

	tests := []struct {
		name     string
		rover    string
		camera   string
		sol      string
		expected int
	}{
		{"Curiosity_FHAZ", "curiosity", "fhaz", data.TestSolWithPhotos, 1},
		{"Opportunity_FHAZ", "opportunity", "fhaz", data.TestSolWithoutPhotos, 0},
		{"Spirit_FHAZ", "spirit", "fhaz", data.TestSolWithoutPhotos, 0},
		{"Curiosity_RHAZ", "curiosity", "rhaz", data.TestSolWithPhotos, 1},
		{"Opportunity_RHAZ", "opportunity", "rhaz", data.TestSolWithoutPhotos, 0},
		{"Spirit_RHAZ", "spirit", "rhaz", data.TestSolWithoutPhotos, 0},
		{"Curiosity_MAST", "curiosity", "mast", data.TestSolWithPhotos, 1},
		{"Curiosity_CHEMCAM", "curiosity", "chemcam", data.TestSolWithPhotos, 1},
		{"Curiosity_MAHLI", "curiosity", "mahli", data.TestSolWithoutPhotos, 0},
		{"Curiosity_MARDI", "curiosity", "mardi", data.TestSolWithoutPhotos, 0},
		{"Curiosity_NAVCAM", "curiosity", "navcam", data.TestSolWithPhotos, 1},
		{"Opportunity_NAVCAM", "opportunity", "navcam", data.TestSolWithPhotos, 1},
		{"Spirit_NAVCAM", "spirit", "navcam", data.TestSolWithPhotos, 1},
		{"Opportunity_PANCAM", "opportunity", "pancam", data.TestSolWithPhotos, 1},
		{"Spirit_PANCAM", "spirit", "pancam", data.TestSolWithPhotos, 1},
		{"Opportunity_MINITES", "opportunity", "minites", data.TestSolWithoutPhotos, 0},
		{"Spirit_MINITES", "spirit", "minites", data.TestSolWithoutPhotos, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := api.GetMarsPhotos(tt.rover, tt.camera, tt.sol)

			assert.NoError(t, err)
			assert.NotNil(t, result)
			if tt.expected > 0 {
				assert.Greater(t, len(result.Photos), 0, "expected photos for "+tt.name)
			} else {
				assert.Empty(t, result.Photos, "expected no photos for "+tt.name)
			}
		})
	}
}

func TestGetMarsPhotosByRover(t *testing.T) {
	SetupRealAPIKey(t)

	rovers := []string{"curiosity", "opportunity", "spirit"}
	for _, rover := range rovers {
		rover := rover
		t.Run("Rover: "+rover, func(t *testing.T) {
			testMarsPhotos(t, rover, "navcam", data.TestSolWithPhotos)
		})
	}
}

func TestGetMarsPhotosByCamera(t *testing.T) {
	SetupRealAPIKey(t)

	cameras := []string{"fhaz", "rhaz", "mast", "chemcam", "navcam"}
	for _, camera := range cameras {
		camera := camera
		t.Run("Camera: "+camera, func(t *testing.T) {
			testMarsPhotos(t, "curiosity", camera, data.TestSolWithPhotos)
		})
	}
}

func TestGetMarsPhotosInvalidRoverReturnsError(t *testing.T) {
	SetupRealAPIKey(t)

	result, err := api.GetMarsPhotos("NonExistingRover", "fhaz", data.TestSolWithPhotos)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetMarsPhotosInvalidCameraReturnsEmptyList(t *testing.T) {
	SetupRealAPIKey(t)

	result, err := api.GetMarsPhotos("opportunity", "NonExistingCamera", data.TestSolWithPhotos)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Empty(t, result.Photos, "expected no photos for ", result)
}

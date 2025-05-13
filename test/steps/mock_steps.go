package steps

import (
	"encoding/json"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/slazarska/mars-go-tests/internal/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func PrepareMockServer(t provider.T, filename string) (*httptest.Server, string, error) {
	var mockResp models.RoverResponse
	err := LoadMockJSON(t, filename, &mockResp)
	if err != nil {
		return nil, "", err
	}

	server := startMockServer(mockResp)
	mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"
	return server, mockBaseURL, nil
}

func startMockServer(mockResp models.RoverResponse) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(mockResp)
	})
	server := httptest.NewServer(handler)
	return server
}

func LoadMockJSON(t provider.T, filename string, target interface{}) error {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	path := filepath.Join(dir, "testdata", filename)
	t.Logf("loading mock file from: %s", path)

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read mock file %s: %v", filename, err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		return fmt.Errorf("failed to unmarshal mock JSON: %v", err)
	}

	return nil
}

func InternalServerErrorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
	})
}

func CheckSuccessfulResponse(t provider.T, resp *models.RoverResponse, err error) {
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func CheckEmptyPhotos(t provider.T, resp *models.RoverResponse) {
	assert.Empty(t, resp.Photos)
}

func CheckOnePhoto(t provider.T, resp *models.RoverResponse) {
	assert.Equal(t, 1, len(resp.Photos))
	assert.Equal(t, "Curiosity", resp.Photos[0].Rover.Name)
	assert.Equal(t, "http://example.com/image.jpg", resp.Photos[0].ImgSrc)
}

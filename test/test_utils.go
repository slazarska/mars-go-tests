package test

import (
	"encoding/json"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/config"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func SetupRealAPIKey(t *testing.T) {
	t.Helper()

	err := config.LoadConfig()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}
	api.SetAPIKey(config.APIKey())
}

func SetTestAPIKey() {
	config.SetAPIKey("test_key")
}

func testMarsPhotos(t *testing.T, rover, camera, sol string) {
	t.Helper()
	result, err := api.GetMarsPhotos(rover, camera, sol)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Greater(t, len(result.Photos), 0, "expected photos, but got none for "+rover+" and "+camera)
}

func createMockServer(t *testing.T, statusCode int, response interface{}, validateQuery bool) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if validateQuery {
			query := r.URL.Query()
			assert.Equal(t, "1000", query.Get("sol"))
			assert.Equal(t, "fhaz", query.Get("camera"))
			assert.Equal(t, "test_key", query.Get("api_key"))
		}

		w.WriteHeader(statusCode)
		if response != nil {
			if err := json.NewEncoder(w).Encode(response); err != nil {
				t.Fatalf("failed to encode mock response: %v", err)
			}
		}
	}))
}

func LoadMockJSON(t *testing.T, filename string, target interface{}) {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working directory: %v", err)
	}

	path := filepath.Join(dir, "testdata", filename)
	t.Logf("Loading mock file from: %s", path) // Логирование пути

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read mock file %s: %v", filename, err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		t.Fatalf("failed to unmarshal mock JSON: %v", err)
	}
}

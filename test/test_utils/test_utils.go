package test_utils

import (
	"encoding/json"
	"fmt"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/config"
	"github.com/slazarska/mars-go-tests/internal/models"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func SetupRealAPIKey(t provider.T) {
	t.WithNewStep("Load configuration", func(sCtx provider.StepCtx) {
		err := config.LoadConfig()
		if err != nil {
			sCtx.Errorf("failed to load config: %v", err)
			sCtx.FailNow()
			return
		}

		apiKey := config.APIKey()
		sCtx.WithNewStep("Set API key", func(sCtx provider.StepCtx) {
			api.SetAPIKey(apiKey)
			sCtx.Log("API key was set successfully")
		})
	})
}

func SetTestAPIKey() {
	config.SetAPIKey("test_key")
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

func CreateMockServer(mockResponse models.RoverResponse) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockResponse)
	}))
}

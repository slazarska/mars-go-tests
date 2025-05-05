package test

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/slazarska/mars-go-tests/test/utils"
	"testing"
)

func TestMockGetMarsPhotos(t *testing.T) {
	runner.Run(t, "MockGetMarsPhotos", func(t provider.T) {
		utils.SetTestAPIKey()

		t.Run("Mock-test with one photo", func(t provider.T) {
			utils.RunTestWithMockData(t, "mock_response.json", 1, "Curiosity", "http://example.com/image.jpg")
		})

		t.Run("Mock-test with empty photo list", func(t provider.T) {
			utils.RunTestWithMockData(t, "empty_response.json", 0, "", "")
		})

		t.Run("Mock-test with internal server error", func(t provider.T) {
			utils.RunTestWithError(t)
		})
	})
}

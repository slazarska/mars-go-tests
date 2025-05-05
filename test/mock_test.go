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

		t.Run("success with one photo", func(t provider.T) {
			utils.RunTestWithMockData(t, "mock_response.json", 1, "Curiosity", "http://example.com/image.jpg")
		})

		t.Run("success with empty photo list", func(t provider.T) {
			utils.RunTestWithMockData(t, "empty_response.json", 0, "", "")
		})

		t.Run("internal server error", func(t provider.T) {
			utils.RunTestWithError(t)
		})
	})
}

/*
func TestMockGetMarsPhotos(t *testing.T) {
	runner.Run(t, "MockGetMarsPhotos", func(t provider.T) {
		SetTestAPIKey()

		t.Run("success with one photo", func(t provider.T) {
			t.Epic("API Tests")
			t.Feature("Mars Photos API")
			t.Story("Get photos with single result")

			var mockResponse models.RoverResponse

			t.WithNewStep("Load mock data", func(ctx provider.StepCtx) {
				tempT := &testing.T{}
				LoadMockJSON(tempT, "mock_response.json", &mockResponse)
				if tempT.Failed() {
					ctx.Errorf("Failed to load mock data")
					ctx.FailNow()
				}

				data, _ := json.Marshal(mockResponse)
				ctx.WithAttachments(allure.NewAttachment(
					"Mock data",
					allure.JSON,
					data,
				))
			})

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(mockResponse)
			}))
			defer server.Close()

			mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"

			result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

			t.WithNewStep("Assertions", func(ctx provider.StepCtx) {
				ctx.Require().NoError(err)
				ctx.Require().NotNil(result)
				ctx.Assert().Equal(1, len(result.Photos))
				ctx.Assert().Equal(12345, result.Photos[0].ID)
				ctx.Assert().Equal("http://example.com/image.jpg", result.Photos[0].ImgSrc)
				ctx.Assert().Equal("Curiosity", result.Photos[0].Rover.Name)
			})
		})

		t.Run("success with empty photo list", func(t provider.T) {
			t.Epic("API Tests")
			t.Feature("Mars Photos API")
			t.Story("Get photos with no results")

			mockResponse := models.RoverResponse{
				Photos: []models.Photo{},
			}

			t.WithNewStep("Load empty mock data", func(ctx provider.StepCtx) {
				data, _ := json.Marshal(mockResponse)
				ctx.WithAttachments(allure.NewAttachment(
					"Mock Empty Data",
					allure.JSON,
					data,
				))
			})

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(mockResponse)
			}))
			defer server.Close()

			mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"

			result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

			t.WithNewStep("Assertions for empty photo list", func(ctx provider.StepCtx) {
				ctx.Require().NoError(err)
				ctx.Require().NotNil(result)
				ctx.Assert().Empty(result.Photos)
			})
		})

		t.Run("internal server error", func(t provider.T) {
			t.Epic("API Tests")
			t.Feature("Mars Photos API")
			t.Story("Handle server error")

			t.WithNewStep("Simulate server error", func(ctx provider.StepCtx) {
				// Дополнительно можно прикрепить информацию о сервере
				ctx.WithAttachments(allure.NewAttachment(
					"Simulating server error",
					allure.Text,
					[]byte("Internal Server Error"),
				))
			})

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
			}))
			defer server.Close()

			mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"

			result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

			t.WithNewStep("Assertions for error response", func(ctx provider.StepCtx) {
				ctx.Require().Error(err)
				ctx.Require().Nil(result)
			})
		})
	})
}


*/
/*
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

*/

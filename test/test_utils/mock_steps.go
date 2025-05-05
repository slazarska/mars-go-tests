package test_utils

import (
	"encoding/json"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/models"
	"net/http"
	"net/http/httptest"
)

func RunTestWithMockData(t provider.T, filename string, expectedPhotoCount int, expectedRoverName, expectedImgSrc string) {
	t.Epic("API Tests")
	t.Feature("Mars Photos API")
	t.Story("Get photos")

	var mockResponse models.RoverResponse
	err := LoadMockJSON(t, filename, &mockResponse)
	if err != nil {
		t.Errorf("failed to load mock data: %v", err)
		t.FailNow()
	}

	t.WithNewStep("Load mock data", func(ctx provider.StepCtx) {
		data, _ := json.Marshal(mockResponse)
		ctx.WithAttachments(allure.NewAttachment(
			"Mock data",
			allure.JSON,
			data,
		))
	})

	server := CreateMockServer(mockResponse)
	defer server.Close()

	mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"
	result, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

	t.WithNewStep("Assertions", func(ctx provider.StepCtx) {
		ctx.Require().NoError(err)
		ctx.Require().NotNil(result)
		ctx.Assert().Equal(expectedPhotoCount, len(result.Photos))

		if expectedPhotoCount > 0 {
			ctx.Assert().Equal(expectedRoverName, result.Photos[0].Rover.Name)
			ctx.Assert().Equal(expectedImgSrc, result.Photos[0].ImgSrc)
		}
	})
}

func RunTestWithError(t provider.T) {
	t.Epic("API Tests")
	t.Feature("Mars Photos API")
	t.Story("Handle server error")

	t.WithNewStep("Simulate server error", func(ctx provider.StepCtx) {
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
}

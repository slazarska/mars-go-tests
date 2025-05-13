package test

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/test/steps"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestMockWithOnePhoto(t *testing.T) {
	runner.Run(t, "MockWithOnePhoto", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mars Rover's Photos")
		t.Story("Mock-server")
		t.Tags("mock", "mock-test")

		steps.SetTestAPIKey()

		server, mockBaseURL, err := steps.PrepareMockServer(t, "mock_response.json")
		assert.NoError(t, err)
		defer server.Close()

		resp, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

		t.WithNewStep("Check successful response", func(sCtx provider.StepCtx) {
			steps.CheckSuccessfulResponse(t, resp, err)
			steps.AttachResponseBodyJSON(sCtx, resp)
		})
		t.WithNewStep("Check photo data", func(sCtx provider.StepCtx) {
			steps.CheckOnePhoto(t, resp)
			steps.AttachPhotoURLs(sCtx, resp)
		})
	})
}

func TestMockWithEmptyPhotoList(t *testing.T) {
	runner.Run(t, "MockWithEmptyPhotoList", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mars Rover's Photos")
		t.Story("Mock-server")
		t.Tags("mock", "mock-test")

		steps.SetTestAPIKey()

		server, mockBaseURL, err := steps.PrepareMockServer(t, "empty_response.json")
		assert.NoError(t, err)
		defer server.Close()

		resp, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

		t.WithNewStep("Check successful response", func(sCtx provider.StepCtx) {
			steps.CheckSuccessfulResponse(t, resp, err)
			steps.AttachResponseBodyJSON(sCtx, resp)
			steps.CheckNoError(sCtx, err)
		})

		t.WithNewStep("Check the list of photos is empty", func(sCtx provider.StepCtx) {
			steps.CheckEmptyPhotos(t, resp)
		})
	})
}

func TestMockInternalServerError(t *testing.T) {
	runner.Run(t, "MockInternalServerError", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mar Rover's Photos")
		t.Story("Mock-server")
		t.Tags("mock", "mock-test")

		steps.SetTestAPIKey()

		handler := steps.InternalServerErrorHandler()
		server := httptest.NewServer(handler)
		defer server.Close()

		mockBaseURL := server.URL + "/rovers/%s/photos?sol=%s&camera=%s&api_key=%s"
		resp, err := api.GetMarsPhotos("curiosity", "fhaz", "1000", mockBaseURL)

		t.WithNewStep("Check Internal Server Error", func(sCtx provider.StepCtx) {
			assert.Error(t, err)
			steps.AttachErrorMessage(sCtx, err.Error())
			assert.Nil(t, resp)
		})
	})
}

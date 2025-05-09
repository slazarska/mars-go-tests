package test

import (
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/test/testdata"
	"github.com/slazarska/mars-go-tests/test/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMarsPhotos(t *testing.T) {
	runner.Run(t, "Get photos by Curiosity and camera", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mars Rover's Photos")
		t.Story("Mars Rover Curiosity's cameras")
		t.Tags("Curiosity", "Mars", "API test", "Integration test")
		t.Severity(allure.BLOCKER)

		utils.SetupRealAPIKey(t)

		testSol := testdata.GetRandomSolCuriosity()

		tests := []struct {
			name   string
			rover  string
			camera string
			sol    string
		}{
			{"Photos for rover Curiosity with camera FHAZ", "curiosity", "fhaz", testSol},
			{"Photos for rover Curiosity with camera RHAZ", "curiosity", "rhaz", testSol},
			{"Photos for rover Curiosity with camera MAST", "curiosity", "mast", testSol},
			{"Photos for rover Curiosity with camera CHEMCAM", "curiosity", "chemcam", testSol},
			{"Photos for rover Curiosity with camera MAHLI", "curiosity", "mahli", testSol},
			{"Photos for rover Curiosity with camera NAVCAM", "curiosity", "navcam", testSol},
			{"Photos for rover Curiosity with camera MARDI", "curiosity", "mardi", testSol},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t provider.T) {
				t.Title(tt.name)
				t.Descriptionf("Test getting photos for Curiosity with camera %s on sol 3466", tt.camera)

				t.WithParameters(
					allure.NewParameter("Rover", tt.rover),
					allure.NewParameter("Camera", tt.camera),
					allure.NewParameter("Sol", tt.sol),
				)

				resp, err := api.GetMarsPhotos(tt.rover, tt.camera, tt.sol)

				t.WithNewStep("Assertions", func(sCtx provider.StepCtx) {
					utils.AssertsGetMarsPhotos(sCtx, tt.name, resp, err, tt.sol, tt.camera)
				})

				t.WithNewStep("Attach additional info", func(sCtx provider.StepCtx) {
					utils.AllureAttachments(sCtx, resp)
				})
			})
		}
	})
}

func TestGetMarsPhotosInvalidRoverReturnsError(t *testing.T) {
	runner.Run(t, "Try to get photos by invalid rover's name returns error", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mars Rover's Photos")
		t.Story("Invalid Rover")
		t.Tags("Error", "Mars", "API test", "Integration test")
		t.Severity(allure.MINOR)

		utils.SetupRealAPIKey(t)

		testSol := testdata.GetRandomSolCuriosity()

		resp, err := api.GetMarsPhotos("NonExistingRover", "fhaz", testSol)

		t.WithNewStep("Check error is returned", func(sCtx provider.StepCtx) {
			assert.Error(sCtx, err)
			errText := err.Error()
			assert.Contains(sCtx, errText, "unexpected status code: 400")
			assert.Contains(sCtx, errText, `"errors":"Invalid Rover Name"`)
			sCtx.WithNewAttachment("Error message", allure.Text, []byte(errText))
			assert.Nil(sCtx, resp)
			sCtx.WithNewAttachment("Raw response", allure.Text, []byte("response is nil"))
		})
	})
}

func TestGetMarsPhotosInvalidCameraReturnsEmptyList(t *testing.T) {
	runner.Run(t, "Try to get photos by invalid camera's name returns the empty list without error", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mars Rover's Photos")
		t.Story("Invalid Camera")
		t.Tags("Error", "Mars", "API test", "Integration test")
		t.Severity(allure.MINOR)

		utils.SetupRealAPIKey(t)

		testSol := testdata.GetRandomSolCuriosity()

		resp, err := api.GetMarsPhotos("opportunity", "NonExistingCamera", testSol)

		t.WithNewStep("Check error", func(sCtx provider.StepCtx) {
			assert.NoError(sCtx, err)
			sCtx.WithNewAttachment("Error Check Result", allure.Text, []byte("No error returned from API"))
		})

		t.WithNewStep("Check response", func(sCtx provider.StepCtx) {
			assert.NotNil(sCtx, resp)
			sCtx.WithNewAttachment("Response Check Result", allure.Text, []byte("Response object is not nil"))
		})

		t.WithNewStep("Check photos empty", func(sCtx provider.StepCtx) {
			assert.Empty(sCtx, resp.Photos, fmt.Sprintf("Expected no photos for invalid camera on sol %s", testSol))
			sCtx.WithNewAttachment(
				"Photos Count",
				allure.Text,
				[]byte(fmt.Sprintf("Received %d photos (expected 0)", len(resp.Photos))),
			)
		})
	})
}

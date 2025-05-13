package test

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/config"
	"github.com/slazarska/mars-go-tests/internal/log"
	"github.com/slazarska/mars-go-tests/test/steps"
	"github.com/slazarska/mars-go-tests/test/testdata"
	"testing"
)

func TestGetMarsPhotosForRandomSol(t *testing.T) {
	runner.Run(t, "Get photos by Curiosity and camera", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mars Rover's Photos")
		t.Story("Mars Rover Curiosity's cameras")
		t.Tags("Curiosity", "Mars", "API test", "Integration test")
		t.Severity(allure.BLOCKER)

		if err := config.SetupRealAPIKey(); err != nil {
			log.Error("failed to setup real API key", "error", err)
			return
		}

		randomSol := testdata.GetRandomSolCuriosity()

		tests := []struct {
			name   string
			rover  string
			camera string
			sol    string
		}{
			{"Photos for rover Curiosity with camera FHAZ", "curiosity", "fhaz", randomSol},
			{"Photos for rover Curiosity with camera RHAZ", "curiosity", "rhaz", randomSol},
			{"Photos for rover Curiosity with camera MAST", "curiosity", "mast", randomSol},
			{"Photos for rover Curiosity with camera CHEMCAM", "curiosity", "chemcam", randomSol},
			{"Photos for rover Curiosity with camera MAHLI", "curiosity", "mahli", randomSol},
			{"Photos for rover Curiosity with camera NAVCAM", "curiosity", "navcam", randomSol},
			{"Photos for rover Curiosity with camera MARDI", "curiosity", "mardi", randomSol},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t provider.T) {
				t.Title(tt.name)
				t.Descriptionf("Test getting photos for Curiosity with camera %s on random sol", tt.camera)

				t.WithParameters(
					allure.NewParameter("Rover", tt.rover),
					allure.NewParameter("Camera", tt.camera),
					allure.NewParameter("Sol", tt.sol),
				)

				resp, err := api.GetMarsPhotos(tt.rover, tt.camera, tt.sol)

				t.WithNewStep("Assertions", func(sCtx provider.StepCtx) {
					steps.CheckError(sCtx, err, tt.sol, tt.camera)
					steps.CheckResponse(sCtx, resp, tt.sol, tt.camera)
					steps.LogPhotoCount(sCtx, resp, tt.sol, tt.camera)
				})

				t.WithNewStep("Attach additional info", func(sCtx provider.StepCtx) {
					steps.AttachResponseBodyJSON(sCtx, resp)
					steps.AttachPhotoURLs(sCtx, resp)
					steps.AttachFirstPhoto(sCtx, resp)
				})
			})
		}
	})
}

func TestGetMarsPhotosForCurrentSol(t *testing.T) {
	runner.Run(t, "Get photos by Curiosity and camera", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mars Rover's Photos")
		t.Story("Mars Rover Curiosity's cameras")
		t.Tags("Curiosity", "Mars", "API test", "Integration test")
		t.Severity(allure.BLOCKER)

		if err := config.SetupRealAPIKey(); err != nil {
			log.Error("failed to setup real API key", "error", err)
			return
		}

		currentSol := testdata.GetRandomSolCuriosity()

		tests := []struct {
			name   string
			rover  string
			camera string
			sol    string
		}{
			{"Photos for rover Curiosity with camera FHAZ", "curiosity", "fhaz", currentSol},
			{"Photos for rover Curiosity with camera RHAZ", "curiosity", "rhaz", currentSol},
			{"Photos for rover Curiosity with camera MAST", "curiosity", "mast", currentSol},
			{"Photos for rover Curiosity with camera CHEMCAM", "curiosity", "chemcam", currentSol},
			{"Photos for rover Curiosity with camera MAHLI", "curiosity", "mahli", currentSol},
			{"Photos for rover Curiosity with camera NAVCAM", "curiosity", "navcam", currentSol},
			{"Photos for rover Curiosity with camera MARDI", "curiosity", "mardi", currentSol},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t provider.T) {
				t.Title(tt.name)
				t.Descriptionf("Test getting photos for Curiosity with camera %s for current sol", tt.camera)

				t.WithParameters(
					allure.NewParameter("Rover", tt.rover),
					allure.NewParameter("Camera", tt.camera),
					allure.NewParameter("Sol", tt.sol),
				)

				resp, err := api.GetMarsPhotos(tt.rover, tt.camera, tt.sol)

				t.WithNewStep("Assertions", func(sCtx provider.StepCtx) {
					steps.CheckError(sCtx, err, tt.sol, tt.camera)
					steps.CheckResponse(sCtx, resp, tt.sol, tt.camera)
					steps.LogPhotoCount(sCtx, resp, tt.sol, tt.camera)
				})

				t.WithNewStep("Attach additional info", func(sCtx provider.StepCtx) {
					steps.AttachResponseBodyJSON(sCtx, resp)
					steps.AttachPhotoURLs(sCtx, resp)
					steps.AttachFirstPhoto(sCtx, resp)
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

		if err := config.SetupRealAPIKey(); err != nil {
			log.Error("failed to setup real API key", "error", err)
			return
		}

		testSol := testdata.GetRandomSolCuriosity()

		_, err := api.GetMarsPhotos("NonExistingRover", "fhaz", testSol)

		t.WithNewStep("Check error is returned", func(sCtx provider.StepCtx) {
			steps.CheckErrorMessage(sCtx, err, []string{
				"unexpected status code: 400",
				`"errors":"Invalid Rover Name"`,
			})
			steps.AttachErrorMessage(sCtx, err.Error())
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

		if err := config.SetupRealAPIKey(); err != nil {
			log.Error("failed to setup real API key", "error", err)
			return
		}

		testSol := testdata.GetRandomSolCuriosity()

		resp, err := api.GetMarsPhotos("opportunity", "NonExistingCamera", testSol)

		t.WithNewStep("Assertions", func(sCtx provider.StepCtx) {
			steps.CheckNoError(sCtx, err)
			steps.CheckNonNilResponse(sCtx, resp)
			steps.CheckPhotosEmpty(sCtx, resp, testSol)
		})
	})
}

package test

import (
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/slazarska/mars-go-tests/internal/api"
	test "github.com/slazarska/mars-go-tests/internal/constants"
	"github.com/slazarska/mars-go-tests/test/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMarsPhotosSpirit(t *testing.T) {
	runner.Run(t, "Get photos by Spirit and camera", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mar Rover's Photos")
		t.Story("Mars Rover Spirit's cameras")
		t.Tags("Spirit", "Mars", "API test", "Integration test")
		t.Severity(allure.BLOCKER)

		test_utils.SetupRealAPIKey(t)

		tests := []struct {
			name   string
			rover  string
			camera string
			sol    string
		}{
			{"Photos for rover Spirit with camera FHAZ", "spirit", "fhaz", test.Sol},
			{"Photos for rover Spirit with camera RHAZ", "spirit", "rhaz", test.Sol},
			{"Photos for rover Spirit with camera NAVCAM", "spirit", "navcam", test.Sol},
			{"Photos for rover Spirit with camera PANCAM", "spirit", "pancam", test.Sol},
			{"Photos for rover Spirit with camera MINITES", "spirit", "minites", test.Sol},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t provider.T) {
				t.Title(tt.name)
				t.Descriptionf("Test getting photos for Spirit with camera %s on sol 54", tt.camera)

				t.WithParameters(
					allure.NewParameter("Rover", tt.rover),
					allure.NewParameter("Camera", tt.camera),
					allure.NewParameter("Sol", tt.sol),
				)

				resp, err := api.GetMarsPhotos(tt.rover, tt.camera, tt.sol)

				t.WithNewStep("Assertions", func(sCtx provider.StepCtx) {
					test_utils.AssertsGetMarsPhotos(sCtx, tt.name, resp, err, tt.sol, tt.camera)
				})

				t.WithNewStep("Attach additional info", func(sCtx provider.StepCtx) {
					test_utils.AllureAttachments(sCtx, resp)
				})
			})
		}
	})
}

func TestGetMarsPhotosOpportunity(t *testing.T) {
	runner.Run(t, "Get photos by Opportunity and camera", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mar Rover's Photos")
		t.Story("Mars Rover Opportunity's cameras")
		t.Tags("Opportunity", "Mars", "API test", "Integration test")
		t.Severity(allure.BLOCKER)

		test_utils.SetupRealAPIKey(t)

		tests := []struct {
			name   string
			rover  string
			camera string
			sol    string
		}{
			{"Photos for rover Opportunity with camera FHAZ", "opportunity", "fhaz", test.Sol},
			{"Photos for rover Opportunity with camera RHAZ", "opportunity", "rhaz", test.Sol},
			{"Photos for rover Opportunity with camera NAVCAM", "opportunity", "navcam", test.Sol},
			{"Photos for rover Opportunity with camera PANCAM", "opportunity", "pancam", test.Sol},
			{"Photos for rover Opportunity with camera MINITES", "opportunity", "minites", test.Sol},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t provider.T) {
				t.Title(tt.name)
				t.Descriptionf("Test getting photos for Opportunity with camera %s on sol 100", tt.camera)

				t.WithParameters(
					allure.NewParameter("Rover", tt.rover),
					allure.NewParameter("Camera", tt.camera),
					allure.NewParameter("Sol", tt.sol),
				)

				resp, err := api.GetMarsPhotos(tt.rover, tt.camera, tt.sol)

				t.WithNewStep("Assertions", func(sCtx provider.StepCtx) {
					test_utils.AssertsGetMarsPhotos(sCtx, tt.name, resp, err, tt.sol, tt.camera)
				})

				t.WithNewStep("Attach additional info", func(sCtx provider.StepCtx) {
					test_utils.AllureAttachments(sCtx, resp)
				})
			})
		}
	})
}

func TestGetMarsPhotosCuriosity(t *testing.T) {
	runner.Run(t, "Get photos by Curiosity and camera", func(t provider.T) {
		t.Epic("Mars Open API")
		t.Feature("Mar Rover's Photos")
		t.Story("Mars Rover Curiosity's cameras")
		t.Tags("Curiosity", "Mars", "API test", "Integration test")
		t.Severity(allure.BLOCKER)

		test_utils.SetupRealAPIKey(t)

		tests := []struct {
			name   string
			rover  string
			camera string
			sol    string
		}{
			{"Photos for rover Curiosity with camera FHAZ", "curiosity", "fhaz", test.Sol},
			{"Photos for rover Curiosity with camera RHAZ", "curiosity", "rhaz", test.Sol},
			{"Photos for rover Curiosity with camera MAST", "curiosity", "mast", test.Sol},
			{"Photos for rover Curiosity with camera CHEMCAM", "curiosity", "chemcam", test.Sol},
			{"Photos for rover Curiosity with camera MAHLI", "curiosity", "mahli", test.Sol},
			{"Photos for rover Curiosity with camera NAVCAM", "curiosity", "navcam", test.Sol},
			{"Photos for rover Curiosity with camera MARDI", "curiosity", "mardi", test.Sol},
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
					test_utils.AssertsGetMarsPhotos(sCtx, tt.name, resp, err, tt.sol, tt.camera)
				})

				t.WithNewStep("Attach additional info", func(sCtx provider.StepCtx) {
					test_utils.AllureAttachments(sCtx, resp)
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

		test_utils.SetupRealAPIKey(t)

		resp, err := api.GetMarsPhotos("NonExistingRover", "fhaz", test.Sol)

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

		test_utils.SetupRealAPIKey(t)

		resp, err := api.GetMarsPhotos("opportunity", "NonExistingCamera", test.Sol)

		t.WithNewStep("Check error", func(sCtx provider.StepCtx) {
			// 1) Ошибки не должно быть
			assert.NoError(sCtx, err)
			sCtx.WithNewAttachment("Error Check Result", allure.Text, []byte("No error returned from API"))
		})

		t.WithNewStep("Check response", func(sCtx provider.StepCtx) {
			// 2) Ответ не nil
			assert.NotNil(sCtx, resp)
			sCtx.WithNewAttachment("Response Check Result", allure.Text, []byte("Response object is not nil"))
		})

		t.WithNewStep("Check photos empty", func(sCtx provider.StepCtx) {
			// 3) Список фото должен быть пустым
			assert.Empty(sCtx, resp.Photos, fmt.Sprintf("Expected no photos for invalid camera on sol %s", test.Sol))
			sCtx.WithNewAttachment(
				"Photos Count",
				allure.Text,
				[]byte(fmt.Sprintf("Received %d photos (expected 0)", len(resp.Photos))),
			)
		})
	})
}

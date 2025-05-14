package steps

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/slazarska/mars-go-tests/internal/log"
	"github.com/slazarska/mars-go-tests/internal/models"
	"github.com/stretchr/testify/assert"
)

func CheckError(sCtx provider.StepCtx, err error, sol, camera string) {
	sCtx.WithNewStep("Check error", func(sCtx provider.StepCtx) {
		assert.NoError(sCtx, err)
		sCtx.WithNewAttachment("Error Check Result", allure.Text, []byte("No error returned from API"))
		log.Info("No error returned", "camera", camera, "sol", sol)
	})
}

func CheckResponse(sCtx provider.StepCtx, resp *models.RoverResponse, sol, camera string) {
	sCtx.WithNewStep("Check response", func(sCtx provider.StepCtx) {
		assert.NotNil(sCtx, resp)
		sCtx.WithNewAttachment("Response Check Result", allure.Text, []byte("Response object is not nil"))
		log.Info("Response is valid", "camera", camera, "sol", sol)
	})
}

func CheckErrorMessage(sCtx provider.StepCtx, err error, expectedMessages []string) {
	sCtx.WithNewStep("Check error and message contents", func(sCtx provider.StepCtx) {
		assert.Error(sCtx, err)

		errText := err.Error()
		for _, msg := range expectedMessages {
			assert.Contains(sCtx, errText, msg)
		}
	})
}

func CheckNoError(sCtx provider.StepCtx, err error) {
	sCtx.WithNewStep("Check no error", func(sCtx provider.StepCtx) {
		assert.NoError(sCtx, err)
		sCtx.WithNewAttachment("Error Check Result", allure.Text, []byte("No error returned from API"))
	})
}

func CheckNonNilResponse(sCtx provider.StepCtx, resp *models.RoverResponse) {
	sCtx.WithNewStep("Check response is not nil", func(sCtx provider.StepCtx) {
		assert.NotNil(sCtx, resp)
		sCtx.WithNewAttachment("Response Check Result", allure.Text, []byte("Response object is not nil"))
	})
}

func CheckPhotosEmpty(sCtx provider.StepCtx, resp *models.RoverResponse, sol string) {
	sCtx.WithNewStep("Check photos list is empty", func(sCtx provider.StepCtx) {
		assert.Empty(sCtx, resp.Photos, fmt.Sprintf("Expected no photos for invalid camera on sol %s", sol))
		sCtx.WithNewAttachment(
			"Photos Count",
			allure.Text,
			[]byte(fmt.Sprintf("Received %d photos (expected 0)", len(resp.Photos))),
		)
	})
}

func LogPhotoCount(sCtx provider.StepCtx, resp *models.RoverResponse, sol, camera string) {
	sCtx.WithNewStep("Log photos count", func(sCtx provider.StepCtx) {
		count := len(resp.Photos)
		var message string

		if count > 0 {
			message = "Received photos"
		} else {
			message = "No photos received"
		}

		logMsg := message + " from camera " + camera + " on sol " + sol
		sCtx.WithNewAttachment("Photos Count Log", allure.Text, []byte(logMsg))
		log.Info(message, "photo_count", count, "camera", camera, "sol", sol)
	})
}

func AttachResponseBodyJSON(sCtx provider.StepCtx, resp *models.RoverResponse) {
	sCtx.WithNewStep("Attach response JSON", func(sCtx provider.StepCtx) {
		rawBody, err := json.MarshalIndent(resp, "", "  ")
		if err == nil {
			sCtx.WithAttachments(allure.NewAttachment("Response JSON", allure.Text, rawBody))
			log.Info("Attached response JSON", "photo_count", len(resp.Photos))
		} else {
			log.Error("Failed to marshal response JSON", "error", err)
		}
	})
}

func AttachPhotoURLs(sCtx provider.StepCtx, resp *models.RoverResponse) {
	sCtx.WithNewStep("Attach photo URLs", func(sCtx provider.StepCtx) {
		if len(resp.Photos) == 0 {
			log.Info("No photo URLs to attach")
			return
		}

		var urls []string
		for _, photo := range resp.Photos {
			urls = append(urls, photo.ImgSrc)
		}
		photoList := strings.Join(urls, "\n")
		sCtx.WithAttachments(allure.NewAttachment("Photo URLs", allure.Text, []byte(photoList)))
		log.Info("Attached photo URLs", "count", len(urls))
	})
}

func AttachFirstPhoto(sCtx provider.StepCtx, resp *models.RoverResponse) {
	sCtx.WithNewStep("Attach first photo", func(sCtx provider.StepCtx) {
		if len(resp.Photos) == 0 {
			log.Info("No photos to attach")
			return
		}

		imgURL := resp.Photos[0].ImgSrc
		respImg, err := http.Get(imgURL)
		if err != nil {
			log.Error("Failed to fetch image", "url", imgURL, "error", err)
			return
		}
		defer respImg.Body.Close()

		imgData, err := io.ReadAll(respImg.Body)
		if err != nil {
			log.Error("Failed to read image data", "url", imgURL, "error", err)
			return
		}

		sCtx.WithAttachments(allure.NewAttachment("First Photo", allure.Jpg, imgData))
		log.Info("Attached first photo", "url", imgURL)
	})
}

func AttachErrorMessage(sCtx provider.StepCtx, errText string) {
	sCtx.WithNewStep("Attach error message", func(sCtx provider.StepCtx) {
		sCtx.WithAttachments(allure.NewAttachment("Error message", allure.Text, []byte(errText)))
		log.Info("Attached error message", "error", errText)
	})
}

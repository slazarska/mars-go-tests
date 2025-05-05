package test_utils

import (
	"encoding/json"
	"fmt"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"

	"github.com/slazarska/mars-go-tests/internal/models"
	"github.com/stretchr/testify/assert"

	"io"
	"net/http"
	"strings"
)

func AssertsGetMarsPhotos(sCtx provider.StepCtx, _ string, resp *models.RoverResponse, err error, sol string, camera string) {
	sCtx.WithNewStep("Check error", func(sCtx provider.StepCtx) {
		assert.NoError(sCtx, err)
		sCtx.WithNewAttachment("Error Check Result", allure.Text, []byte("No error returned from API"))
	})

	sCtx.WithNewStep("Check response", func(sCtx provider.StepCtx) {
		assert.NotNil(sCtx, resp)
		sCtx.WithNewAttachment("Response Check Result", allure.Text, []byte("Response object is not nil"))
	})

	sCtx.WithNewStep("Log photos count", func(sCtx provider.StepCtx) {
		photoCount := len(resp.Photos)
		var logMessage string

		if photoCount > 0 {
			logMessage = fmt.Sprintf("received %d photos from camera %s on sol %s", photoCount, camera, sol)
		} else {
			logMessage = fmt.Sprintf("no photos received from camera %s on sol %s", camera, sol)
		}

		sCtx.WithNewAttachment("Photos Count Log", allure.Text, []byte(logMessage))
	})
}

func AllureAttachments(sCtx provider.StepCtx, resp *models.RoverResponse) {
	rawBody, _ := json.MarshalIndent(resp, "", "  ")
	sCtx.WithAttachments(allure.NewAttachment("Response JSON", allure.Text, rawBody))

	var urls []string
	for _, photo := range resp.Photos {
		urls = append(urls, photo.ImgSrc)
	}
	if len(urls) > 0 {
		photoList := strings.Join(urls, "\n")
		sCtx.WithAttachments(allure.NewAttachment("Photo URLs", allure.Text, []byte(photoList)))
	}

	if len(resp.Photos) > 0 {
		imgURL := resp.Photos[0].ImgSrc
		respImg, err := http.Get(imgURL)
		if err == nil {
			defer respImg.Body.Close()
			imgData, err := io.ReadAll(respImg.Body)
			if err == nil {
				sCtx.WithAttachments(allure.NewAttachment("First Photo", allure.Jpg, imgData))
			}
		}
	}
}

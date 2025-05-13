package steps

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/slazarska/mars-go-tests/internal/api"
	"github.com/slazarska/mars-go-tests/internal/config"
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

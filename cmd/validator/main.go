package main

import (
	"github.com/sirupsen/logrus"

	"maildefender/validator/internal/api"
	"maildefender/validator/internal/configuration"
)

func main() {
	logrus.WithFields(logrus.Fields{
		"engine": configuration.EngineBaseEndpoint(),
	}).Info("Starting validator...")

	api.RegisterHandlers()
	if err := api.Run(); err != nil {
		logrus.WithError(err).Error("API stopped")
	}
}

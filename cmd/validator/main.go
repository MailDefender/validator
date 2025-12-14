package main

import (
	"github.com/sirupsen/logrus"

	"maildefender/validator/internal/api"
)

func main() {
	logrus.Info("Starting validator...")

	api.RegisterHandlers()
	if err := api.Run(); err != nil {
		logrus.WithError(err).Error("API stopped")
	}
}

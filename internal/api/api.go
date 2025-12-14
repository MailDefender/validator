package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"maildefender/validator/internal/api/handlers"
)

var engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	engine.Use(gin.LoggerWithWriter(logrus.New().Writer()))
}

func RegisterHandlers() {
	engine.GET("/validate/:token", handlers.ValidateToken)
}

func Run() error {
	return engine.Run(":8080")
}

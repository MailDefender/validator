package api

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"maildefender/validator/internal/api/handlers"
)

var engine *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)
	engine = gin.New()
	engine.Use(logger())
}

func RegisterHandlers() {
	engine.GET("/validate/:token", handlers.ValidateToken)
}

func Run() error {
	return engine.Run(":8080")
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request

		c.Next()

		// after request
		latency := time.Since(t)

		// access the status we are sending
		status := c.Writer.Status()

		logger := logrus.WithFields(logrus.Fields{
			"status":  status,
			"latency": latency,
			"method":  c.Request.Method,
		})

		logBody := fmt.Sprintf("Handling %s", c.Request.URL.Path)

		switch {
		case status < 300:
			logger.Info(logBody)
			break
		case status >= 300 && status < 400:
			logger.Warn(logBody)
			break
		case status >= 400:
			logger.Error(logBody)
		}
	}
}

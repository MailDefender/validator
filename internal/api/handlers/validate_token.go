package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"maildefender/validator/internal/services"
	"maildefender/validator/internal/templates"
)

type ValidateTokenIn struct {
	Token string `uri:"token" binding:"required"`
}

const (
	expiredToken          string = "expired token"
	invalidToken          string = "invalid token"
	alreadyValidatedToken string = "token already validated"
)

func ValidateToken(c *gin.Context) {
	var in ValidateTokenIn
	if err := c.ShouldBindUri(&in); err != nil {
		logrus.WithError(err).Warn("cannot bind Uri parameters")
		c.Data(http.StatusBadRequest, "text/html", []byte(templates.FailurePage()))
		return
	}

	logger := logrus.WithFields(logrus.Fields{
		"token": in.Token,
	})

	logger.Info("Validating token")

	out, httpCode, err := services.ValidateTokenWithEngine(in.Token)

	if httpCode == http.StatusNotFound || httpCode == http.StatusInternalServerError || err != nil {
		logger.WithError(err).WithFields(logrus.Fields{"http_status": httpCode}).Warn("cannot valid token")
		c.Data(httpCode, "text/html", []byte(templates.FailurePage()))
		return
	}

	var htmlPage string
	switch out.Error {
	case alreadyValidatedToken:
		htmlPage = templates.AlreadyValidatedPage()
		logger.Info("This token is already validated")
	case expiredToken:
		htmlPage = templates.ExpiredPage()
		logger.Info("This token is expired")
	default:
		htmlPage = templates.SuccessPage()
		logger.Info("Token successfully validated")
	}

	c.Data(http.StatusOK, "text/html", []byte(htmlPage))

}

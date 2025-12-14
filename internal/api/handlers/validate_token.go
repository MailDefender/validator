package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		c.Data(http.StatusBadRequest, "text/html", []byte(templates.FailurePage()))
		return
	}

	out, httpCode, err := services.ValidateTokenWithEngine(in.Token)

	if httpCode == http.StatusNotFound || httpCode == http.StatusInternalServerError || err != nil {
		c.Data(httpCode, "text/html", []byte(templates.FailurePage()))
		return
	}

	var htmlPage string
	switch out.Error {
	case alreadyValidatedToken:
		htmlPage = templates.AlreadyValidatedPage()
	case expiredToken:
		htmlPage = templates.ExpiredPage()
	default:
		htmlPage = templates.SuccessPage()
	}

	c.Data(http.StatusOK, "text/html", []byte(htmlPage))

}

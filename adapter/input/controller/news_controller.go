package controller

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/input/model/request"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/logger"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct{}

func NewNewsController() *newsController {
	return &newsController{}
}

func (nc *newsController) GetNews(c *gin.Context) {
	logger.Info("Init GetNews controller api")
	newsRequest := request.NewsRequest{}

	if err := c.ShouldBindQuery(&newsRequest); err != nil {
		logger.Error("Error trying to validate fields from request", err)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return
	}
}

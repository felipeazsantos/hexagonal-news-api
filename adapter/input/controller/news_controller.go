package controller

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/input/model/request"
	"github.com/felipeazsantos/hexagonal-news-api/application/domain"
	"github.com/felipeazsantos/hexagonal-news-api/application/port/input"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/logger"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/validation"
	"github.com/gin-gonic/gin"
)

type newsController struct {
	newsUseCase input.NewsUseCase
}

func NewNewsController(newsUseCase input.NewsUseCase) *newsController {
	return &newsController{
		newsUseCase: newsUseCase,
	}
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

	newsReqDomain := domain.NewsReqDomain{
		Subject: newsRequest.Subject,
		From:    newsRequest.From,
	}

	_, _ = nc.newsUseCase.GetNewsService(newsReqDomain)
}

package routes

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/input/controller"
	"github.com/felipeazsantos/hexagonal-news-api/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsService := service.NewNewsService()
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}

package routes

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/input/controller"
	"github.com/felipeazsantos/hexagonal-news-api/adapter/output/news_http"
	"github.com/felipeazsantos/hexagonal-news-api/application/service"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsClient := news_http.NewNewsClient()
	newsService := service.NewNewsService(newsClient)
	newsController := controller.NewNewsController(newsService)

	r.GET("/news", newsController.GetNews)
}

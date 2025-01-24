package routes

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/input/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	newsController := controller.NewNewsController()

	r.GET("/news", newsController.GetNews)
}

package main

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/input/routes"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Info("About to start application")

	router := gin.Default()
	routes.InitRoutes(router)

	if err := router.Run(":8081"); err != nil {
		logger.Error("Error trying to init application on port :8081", err)
		return
	}
}

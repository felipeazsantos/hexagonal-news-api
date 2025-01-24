package main

import (
	"github.com/felipeazsantos/hexagonal-news-api/adapter/input/routes"
	"github.com/felipeazsantos/hexagonal-news-api/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start application")

	if err := godotenv.Load(); err != nil {
		logger.Error("Error when trying to load env configs", err)
		return
	}

	router := gin.Default()
	routes.InitRoutes(router)

	if err := router.Run(":8081"); err != nil {
		logger.Error("Error trying to init application on port :8081", err)
		return
	}
}

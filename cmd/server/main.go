package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortener/internal/api"
	"github.com/muditsaxena1/url-shortener/internal/config"
	"github.com/muditsaxena1/url-shortener/internal/services"
	"github.com/muditsaxena1/url-shortener/internal/storage"
)

func main() {

	router := gin.Default()

	// Initialize storage and services
	cacheStore := storage.NewInMemoryStorage()
	dbStore := storage.NewMySQLStorage()
	shortenerService := services.NewShortenerService(cacheStore, dbStore)

	// Initialize API routes
	api.SetupRoutes(router, shortenerService)

	config := config.LoadConfig()

	// Start the server
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

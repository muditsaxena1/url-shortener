package main

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortener/internal/api"
	"github.com/muditsaxena1/url-shortener/internal/config"
	"github.com/muditsaxena1/url-shortener/internal/services"
	"github.com/muditsaxena1/url-shortener/internal/storage"
	"github.com/muditsaxena1/url-shortener/internal/utils"
)

func main() {

	router := gin.Default()

	// Initialize storage and services
	cacheStore := storage.NewInMemoryStorage()
	dbStore := storage.NewMySQLStorage()

	var instanceID int64
	var err error
	if val := os.Getenv("INSTANCE_ID"); val == "" {
		panic("INSTANCE_ID not found")
	} else if instanceID, err = strconv.ParseInt(val, 10, 64); err != nil {
		panic("INSTANCE_ID is not a number")
	} else if instanceID < 0 || instanceID > 15 {
		panic("INSTANCE_ID should be between 0 and 15")
	}
	sf := utils.NewSnowflake(instanceID)
	shortenerService := services.NewShortenerService(cacheStore, dbStore, sf)

	// Initialize API routes
	api.SetupRoutes(router, shortenerService)

	config := config.LoadConfig()

	// Start the server
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

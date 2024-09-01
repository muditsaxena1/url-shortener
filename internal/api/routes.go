package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortener/internal/services"
)

func SetupRoutes(router *gin.Engine, shortenerService *services.ShortenerService) {
	h := &Handlers{
		ShortenerService: shortenerService,
	}

	router.POST("/shorten", h.ShortenURL)
	router.GET("/r/:shortCode", h.Redirect)
	router.GET("/metrics/top-domains", h.GetTopDomains)
}

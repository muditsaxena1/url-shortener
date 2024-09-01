package api

import (
	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortener/internal/services"
)

func SetupRoutes(router *gin.Engine, shortenerService *services.ShortenerService) {
	h := &Handlers{
		ShortenerService: shortenerService,
	}

	router.POST("/shorten", h.shortenURL)
	router.GET("/r/:shortCode", h.redirect)
	router.GET("/metrics/top-domains", h.getTopDomains)
}

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muditsaxena1/url-shortener/internal/models"
	"github.com/muditsaxena1/url-shortener/internal/services"
)

type Handlers struct {
	ShortenerService *services.ShortenerService
}

func (h *Handlers) ShortenURL(c *gin.Context) {
	var request models.ShortenRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	shortenedURL, err := h.ShortenerService.ShortenURL(request.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortened_url": shortenedURL})
}

func (h *Handlers) Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	originalURL, err := h.ShortenerService.ResolveURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

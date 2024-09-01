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

func (h *Handlers) shortenURL(c *gin.Context) {
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

func (h *Handlers) redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")
	originalURL, err := h.ShortenerService.ResolveURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func (h *Handlers) getTopDomains(c *gin.Context) {
	topDomains := h.ShortenerService.GetTopDomains()
	c.JSON(http.StatusOK, gin.H{"top_domains": topDomains})
}

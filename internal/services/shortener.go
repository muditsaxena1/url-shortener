package services

import (
	"fmt"
	"os"

	"github.com/muditsaxena1/url-shortener/internal/storage"
	"github.com/muditsaxena1/url-shortener/internal/utils"
)

type ShortenerService struct {
	cache storage.CacheStorage
	db    storage.DatabaseStorage
	sf    *utils.Snowflake
}

func NewShortenerService(cache storage.CacheStorage, db storage.DatabaseStorage, sf *utils.Snowflake) *ShortenerService {
	return &ShortenerService{
		cache: cache,
		db:    db,
		sf:    sf,
	}
}

func (s *ShortenerService) ShortenURL(originalURL string) (string, error) {
	if shortCode, err := s.db.GetShortCode(originalURL); err == nil {
		return os.Getenv("DOMAIN_NAME") + "/r/" + shortCode, nil
	}

	shortCode := s.sf.GenerateShortCode()

	// Save the URL mapping if it doesn't already exist
	if err := s.db.SaveURLMapping(shortCode, originalURL); err != nil {
		return "", err
	}

	return os.Getenv("DOMAIN_NAME") + "/r/" + shortCode, nil
}

func (s *ShortenerService) ResolveURL(shortCode string) (string, error) {
	if originalURL, err := s.cache.GetOriginalURL(shortCode); err == nil {
		return originalURL, err
	}
	originalURL, err := s.db.GetOriginalURL(shortCode)
	if err != nil {
		return "", err
	}
	if err := s.cache.SaveURLMapping(shortCode, originalURL); err != nil {
		fmt.Println("Error while saving data in cache", err.Message)
	}
	return originalURL, nil
}

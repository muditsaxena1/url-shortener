package services

import (
	"fmt"

	"github.com/muditsaxena1/url-shortener/internal/storage"
	"github.com/muditsaxena1/url-shortener/internal/utils"
)

type ShortenerService struct {
	cache storage.CacheStorage
	db    storage.DatabaseStorage
}

func NewShortenerService(cache storage.CacheStorage, db storage.DatabaseStorage) *ShortenerService {
	return &ShortenerService{
		cache: cache,
		db:    db,
	}
}

func (s *ShortenerService) ShortenURL(originalURL string) (string, error) {
	if shortURL, err := s.db.GetShortURL(originalURL); err == nil {
		return shortURL, nil
	}

	shortCode := utils.GenerateShortCode(originalURL)

	// Save the URL mapping if it doesn't already exist
	if err := s.db.SaveURLMapping(shortCode, originalURL); err != nil {
		return "", err
	}

	return "http://localhost:8080/r/" + shortCode, nil
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

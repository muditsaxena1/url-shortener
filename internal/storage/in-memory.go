package storage

import (
	"net/http"
	"sync"

	"github.com/muditsaxena1/url-shortener/internal/errors"
)

type InMemoryStorage struct {
	urlMappings map[string]string
	mu          sync.RWMutex
}

func NewInMemoryStorage() CacheStorage {
	return &InMemoryStorage{urlMappings: make(map[string]string)}
}

func (s *InMemoryStorage) SaveURLMapping(shortCode string, originalURL string) *errors.CustomError {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.urlMappings[shortCode] = originalURL
	return nil
}

func (s *InMemoryStorage) GetOriginalURL(shortCode string) (string, *errors.CustomError) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, exists := s.urlMappings[shortCode]
	if !exists {
		return "", errors.NewCustomError(http.StatusNotFound, "URL not found")
	}
	return url, nil
}

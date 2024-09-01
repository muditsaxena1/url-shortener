package storage

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/muditsaxena1/url-shortener/internal/errors"
)

type MySQLStorage struct {
	urlMappings  map[string]string
	domainCounts map[string]int
	mu           sync.RWMutex
}

func NewMySQLStorage() DatabaseStorage {
	return &MySQLStorage{
		urlMappings:  make(map[string]string),
		domainCounts: make(map[string]int),
	}
}

func (s *MySQLStorage) SaveURLMapping(shortCode string, originalURL string) *errors.CustomError {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.urlMappings[shortCode] = originalURL
	return nil
}

func (s *MySQLStorage) GetOriginalURL(shortCode string) (string, *errors.CustomError) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, exists := s.urlMappings[shortCode]
	if !exists {
		return "", errors.NewCustomError(http.StatusNotFound, "URL not found")
	}
	return url, nil
}

func (s *MySQLStorage) GetShortCode(originalURL string) (string, *errors.CustomError) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for k, v := range s.urlMappings {
		if v == originalURL {
			fmt.Println("Short url already exists for", originalURL)
			return k, nil
		}
	}
	return "", errors.NewCustomError(http.StatusNotFound, "original url not found")
}

func (s *MySQLStorage) IncrementDomainCount(domain string) *errors.CustomError {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.domainCounts[domain]++
	return nil
}

func (s *MySQLStorage) GetDomainCounts() map[string]int {
	copyDomainCounts := make(map[string]int)
	s.mu.RLock()
	defer s.mu.RUnlock()

	for k, v := range s.domainCounts {
		copyDomainCounts[k] = v
	}

	return copyDomainCounts
}

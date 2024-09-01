package storage

import "github.com/muditsaxena1/url-shortener/internal/errors"

type CacheStorage interface {
	SaveURLMapping(shortCode string, originalURL string) *errors.Error
	GetOriginalURL(shortCode string) (string, *errors.Error)
}

type DatabaseStorage interface {
	SaveURLMapping(shortCode string, originalURL string) *errors.Error
	GetOriginalURL(shortCode string) (string, *errors.Error)
	GetShortCode(originalURL string) (string, *errors.Error)
	IncrementDomainCount(domain string) *errors.Error
	GetDomainCounts() map[string]int
}

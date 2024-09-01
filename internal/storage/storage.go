package storage

import "github.com/muditsaxena1/url-shortener/internal/errors"

type CacheStorage interface {
	SaveURLMapping(shortCode string, originalURL string) *errors.CustomError
	GetOriginalURL(shortCode string) (string, *errors.CustomError)
}

type DatabaseStorage interface {
	SaveURLMapping(shortCode string, originalURL string) *errors.CustomError
	GetOriginalURL(shortCode string) (string, *errors.CustomError)
	GetShortCode(originalURL string) (string, *errors.CustomError)
	IncrementDomainCount(domain string) *errors.CustomError
	GetDomainCounts() map[string]int
}

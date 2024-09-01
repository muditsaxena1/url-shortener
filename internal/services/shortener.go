package services

import (
	"fmt"
	"os"
	"strings"

	"github.com/muditsaxena1/url-shortener/internal/models"
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

	// Extract domain and increment domain count
	domain := strings.Split(originalURL, "/")[2]
	s.db.IncrementDomainCount(domain)

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

func (s *ShortenerService) GetTopDomains() []models.Domain {
	domainCounts := s.db.GetDomainCounts()

	var max, secondMax, thirdMax models.Domain

	for k, v := range domainCounts {
		if v >= max.VisitCount {
			thirdMax = secondMax
			secondMax = max
			max = models.Domain{
				DomainURL:  k,
				VisitCount: v,
			}
		} else if v >= secondMax.VisitCount {
			thirdMax = secondMax
			secondMax = models.Domain{
				DomainURL:  k,
				VisitCount: v,
			}
		} else if v >= thirdMax.VisitCount {
			thirdMax = models.Domain{
				DomainURL:  k,
				VisitCount: v,
			}
		}
	}

	return []models.Domain{max, secondMax, thirdMax}
}

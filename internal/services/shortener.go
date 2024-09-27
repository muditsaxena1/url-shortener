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

/*
*
Creates a new instance of ShortenerService with the provided cache storage, database storage, and Snowflake generator.

@param cache The cache storage to be used by the ShortenerService.
@param db The database storage to be used by the ShortenerService.
@param sf The Snowflake generator to be used by the ShortenerService.

@return A pointer to the newly created ShortenerService instance.
*/
func NewShortenerService(cache storage.CacheStorage, db storage.DatabaseStorage, sf *utils.Snowflake) *ShortenerService {
	return &ShortenerService{
		cache: cache,
		db:    db,
		sf:    sf,
	}
}

/**
 * ShortenURL takes an original URL, generates a short code for it, and saves the URL mapping if it doesn't already exist.
 * It also increments the count for the domain of the original URL.
 *
 * @param originalURL The original URL to be shortened.
 * @return The shortened URL or an error if the operation fails.
 */
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

/**
 * ResolveURL resolves the original URL for a given short code.
 *
 * Parameters:
 * - shortCode: the short code for which the original URL needs to be resolved
 *
 * Returns:
 * - string: the original URL corresponding to the short code
 * - error: an error, if any, encountered during the resolution process
 */
func (s *ShortenerService) ResolveURL(shortCode string) (string, error) {
	if originalURL, err := s.cache.GetOriginalURL(shortCode); err == nil {
		return originalURL, err
	}
	originalURL, err := s.db.GetOriginalURL(shortCode)
	if err != nil {
		return "", err
	}
	go func() {
		if err := s.cache.SaveURLMapping(shortCode, originalURL); err != nil {
			fmt.Println("Error while saving data in cache", err.Message)
		}
	}()
	return originalURL, nil
}

/**
 * GetTopDomains returns the top three domains based on visit counts from the database.
 * It iterates through the domain counts, identifies the domains with the highest visit counts,
 * and returns them as a slice of models.Domain structs.
 *
 * @return []models.Domain - A slice containing the top three domains with the highest visit counts.
 */
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

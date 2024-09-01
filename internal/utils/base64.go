package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func GenerateShortCode(url string) string {
	hash := sha256.Sum256([]byte(url))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	return encoded[:8] // Shorten to 8 characters
}

// encodeIDToBase64 converts a 48-bit ID to a Base64 encoded string
func encodeIDToBase64(id int64) string {
	// Create a 6-byte array
	bytes := make([]byte, 6)

	// Convert the 48-bit integer to a 6-byte array
	for i := 5; i >= 0; i-- {
		bytes[i] = byte(id & 0xFF)
		id >>= 8
	}

	// Encode the byte array to a Base64 string
	return base64.RawURLEncoding.EncodeToString(bytes)
}

package shortener

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateShortLink(initialLink, userID string) string {
	urlHash := sha256.Sum256([]byte(initialLink + userID)) // fixed-size array
	hashStr := hex.EncodeToString(urlHash[:])              // → hex string
	if len(hashStr) < 8 {
		return hashStr
	}
	return hashStr[:8]
}

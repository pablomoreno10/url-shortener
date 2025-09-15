package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateShortLink(t *testing.T) {
	userID := "user123"
	a := GenerateShortLink("https://example.com/x", userID)
	b := GenerateShortLink("https://example.com/x", userID)

	assert.Equal(t, a, b, "same input should yield same code")
	assert.Len(t, a, 8, "code should be 8 characters")
}

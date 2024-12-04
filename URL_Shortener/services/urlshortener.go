package services

import (
	"math/rand"
	"sync"
	"time" // This import is necessary for rand.Seed
)

// URLShortener handles URL shortening and retrieval
type URLShortener struct {
	urlStore map[string]string
	mutex    sync.RWMutex
}

// NewURLShortener creates and initializes a new URLShortener instance
func NewURLShortener() *URLShortener {
	rand.Seed(time.Now().UnixNano()) // Ensures unique random values
	return &URLShortener{
		urlStore: make(map[string]string),
	}
}

// GenerateShortCode creates a random short code
func (us *URLShortener) GenerateShortCode() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	shortCode := make([]byte, 6)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortCode)
}

// ShortenURL stores the long URL and returns a short code
func (us *URLShortener) ShortenURL(longURL string) string {
	us.mutex.Lock()
	defer us.mutex.Unlock()

	shortCode := us.GenerateShortCode()
	us.urlStore[shortCode] = longURL
	return shortCode
}

// GetLongURL retrieves the original URL from a short code
func (us *URLShortener) GetLongURL(shortCode string) (string, bool) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	longURL, exists := us.urlStore[shortCode]
	return longURL, exists
}

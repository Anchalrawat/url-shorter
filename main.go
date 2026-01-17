package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

type URLShortener struct {
	urls map[string]string
	mu   sync.RWMutex
}

var shortener = &URLShortener{
	urls: make(map[string]string),
}

// THIS IS THE REPLACEMENT FOR MATH.RAND
func generateShortKey(originalURL string) string {
	// 1. Create a SHA-256 hash of the original URL
	algorithm := sha256.New()
	algorithm.Write([]byte(originalURL))
	hashBytes := algorithm.Sum(nil)

	// 2. Encode the binary hash into a URL-safe string
	// We use URLEncoding to avoid symbols like '+' or '/' that break URLs
	encoded := base64.URLEncoding.EncodeToString(hashBytes)

	// 3. Take the first 8 characters
	// This gives us a short, unique ID like "A1b2C3d4"
	return encoded[:8]
}

func makeShortURL(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	originalURL := req.FormValue("url")
	if originalURL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	// Generate the key using the HASH of the URL input
	shortKey := generateShortKey(originalURL)

	shortener.mu.Lock()
	shortener.urls[shortKey] = originalURL
	shortener.mu.Unlock()

	constructURL := fmt.Sprintf("http://localhost:8090/%s", shortKey)
	fmt.Fprintf(w, "Shortened URL: %s\n", constructURL)
}

func handleRedirect(w http.ResponseWriter, req *http.Request) {
	shortKey := strings.TrimPrefix(req.URL.Path, "/")

	if shortKey == "" {
		http.Error(w, "Short key is missing", http.StatusBadRequest)
		return
	}

	shortener.mu.RLock()
	originalURL, ok := shortener.urls[shortKey]
	shortener.mu.RUnlock()

	if !ok {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, req, originalURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/shorten", makeShortURL)
	http.HandleFunc("/", handleRedirect)

	fmt.Println("Server starting on port :8090...")
	http.ListenAndServe(":8090", nil)
}

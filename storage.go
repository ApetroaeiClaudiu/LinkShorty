package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type URLStore struct {
	mu   sync.RWMutex
	urls map[string]string
	file string
}

func NewURLStore(file string) *URLStore {
	store := &URLStore{
		urls: make(map[string]string),
		file: file,
	}
	store.load() // Load data from the file if it exists
	return store
}

func (s *URLStore) Add(longURL string) string {
	slug := generateSlug(6)
	s.urls[slug] = longURL
	s.save() // Save the updated store
	return slug
}

func (s *URLStore) Get(slug string) (string, bool) {
	url, ok := s.urls[slug]
	return url, ok
}

func (s *URLStore) All() map[string]string {
	copy := make(map[string]string)
	for k, v := range s.urls {
		copy[k] = v
	}
	return copy
}

func (s *URLStore) save() {
	file, err := os.Create(s.file)
	if err != nil {
		fmt.Println("Error creating file:", err)
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(s.urls); err != nil {
		fmt.Println("Error encoding JSON:", err)
		panic(err)
	}
}

func (s *URLStore) load() {
	file, err := os.Open(s.file)
	if err != nil {
		if os.IsNotExist(err) {
			return // No file, nothing to load
		}
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&s.urls); err != nil {
		panic(err)
	}
}

// generateSlug generates a random string of the specified length
func generateSlug(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func (s *URLStore) FindSlugByURL(url string) string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for slug, storedURL := range s.urls {
		if storedURL == url {
			return slug
		}
	}
	return ""
}

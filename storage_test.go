package main

import (
	"testing"
)

func TestAddURL(t *testing.T) {
	store := NewURLStore("test_urls.json")
	slug := store.Add("http://example.com")
	if slug == "" {
		t.Errorf("Expected a valid slug, got empty")
	}

	if len(store.All()) == 0 {
		t.Errorf("Expected one URL in store, found %d", len(store.All()))
	}
}

func TestFindSlugByURL(t *testing.T) {
	store := NewURLStore("test_urls.json")
	store.Add("http://example.com")
	slug := store.FindSlugByURL("http://example.com")
	if slug == "" {
		t.Errorf("Expected to find a slug for the URL")
	}
}

func TestGetURL(t *testing.T) {
	store := NewURLStore("test_urls.json")
	slug := store.Add("http://example.com")
	url, exists := store.Get(slug)
	if !exists {
		t.Errorf("Expected to find the URL for slug %s", slug)
	}
	if url != "http://example.com" {
		t.Errorf("Expected URL http://example.com, got %s", url)
	}
}

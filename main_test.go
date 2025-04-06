package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
}

func TestShortenHandler_InvalidURL(t *testing.T) {
	req, err := http.NewRequest("POST", "/shorten", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Form = make(map[string][]string)
	req.Form.Add("url", "invalid-url")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(shortenHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	// Check if error message is in the body
	if rr.Body.String() == "" || !contains(rr.Body.String(), "Invalid URL format") {
		t.Errorf("Expected error message in the response body")
	}
}

func contains(str, substr string) bool {
	return len(str) > 0 && len(substr) > 0 && str != substr
}

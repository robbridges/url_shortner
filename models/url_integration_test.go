package models

import (
	"echo_url_shortner/data"
	"testing"
)

func TestUrlService_InsertUrl(t *testing.T) {
	cfg := data.DefaultTestPostgresConfig()
	db, err := data.Open(cfg)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}
	urlService := UrlService{DB: db}
	t.Run("Happy path", func(t *testing.T) {
		err = urlService.InsertUrl(&Url{
			Url:      "https://example.com",
			ShortUrl: "https://example.com/short",
			leetcode: false,
		})
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}
		// Get the added url back and check if it's the same
		url, err := urlService.GetUrl("https://example.com/short")
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}
		if url != "https://example.com" {
			t.Errorf("Expected url to be %s, but got %s", "https://example.com", url)
		}
		urlService.DeleteUrl("https://example.com/short")
	})
	t.Run("Duplicate url", func(t *testing.T) {
		err = urlService.InsertUrl(&Url{
			Url:      "https://abc.com",
			ShortUrl: "https://aaa",
			leetcode: false,
		})
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}
		err = urlService.InsertUrl(&Url{
			Url:      "https://abc.com",
			ShortUrl: "https://aaa",
			leetcode: false,
		})
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
		urlService.DeleteUrl("https://aaa")
	})
}

func TestUrlService_GetUrl(t *testing.T) {
	cfg := data.DefaultTestPostgresConfig()
	db, err := data.Open(cfg)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}
	urlService := UrlService{DB: db}
	t.Run("Happy path", func(t *testing.T) {

		url, err := urlService.GetUrl("abcdef")
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}
		if url != "https://www.google.com" {
			t.Errorf("Expected url to be %s, but got %s", "https://example.com", url)
		}
	})
	t.Run("Url not found", func(t *testing.T) {
		_, err := urlService.GetUrl("https://example.com/short2")
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
}

func TestUrlServiceIntegration_DeleteUrl(t *testing.T) {
	cfg := data.DefaultTestPostgresConfig()
	db, err := data.Open(cfg)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}
	urlService := UrlService{DB: db}
	t.Run("Happy path", func(t *testing.T) {
		err = urlService.InsertUrl(&Url{
			Url:      "https://example.com",
			ShortUrl: "https://example.com/short",
			leetcode: false,
		})
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}
		err = urlService.DeleteUrl("https://example.com/short")
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}
		_, err = urlService.GetUrl("https://example.com/short")
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
	t.Run("Url not found", func(t *testing.T) {
		err = urlService.DeleteUrl("https://example.com/short2")
		if err == nil || err.Error() != "URL not found" {
			t.Errorf("Expected 'URL not found' error, but got %v", err)
		}
	})
}

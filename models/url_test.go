package models

import "testing"

func TestUrlServiceMock_InsertUrl(t *testing.T) {
	mockUrlService := &UrlServiceMock{
		DB: []*Url{},
	}

	url := &Url{
		Url:      "https://example.com",
		ShortUrl: "https://example.com/short",
	}
	t.Run("Happy path", func(t *testing.T) {
		mockUrlService.InsertUrl(url)

		if len(mockUrlService.DB) != 1 {
			t.Errorf("Expected DB length to be 1, but got %d", len(mockUrlService.DB))
		}

		if mockUrlService.DB[0].Url != url.Url {
			t.Errorf("Expected url to be %s, but got %s", url.Url, mockUrlService.DB[0].Url)
		}

		if mockUrlService.DB[0].ShortUrl != url.ShortUrl {
			t.Errorf("Expected short url to be %s, but got %s", url.ShortUrl, mockUrlService.DB[0].ShortUrl)
		}
	})
	t.Run("Duplicate url", func(t *testing.T) {
		mockUrlService.InsertUrl(url)

		if len(mockUrlService.DB) != 1 {
			t.Errorf("Expected DB length to be 1, but got %d", len(mockUrlService.DB))
		}
	})
}

func TestUrlServiceMock_GetUrl(t *testing.T) {
	mockUrlService := &UrlServiceMock{
		DB: []*Url{
			{
				Url:      "https://example.com",
				ShortUrl: "https://example.com/short",
			},
		},
	}
	t.Run("Happy path", func(t *testing.T) {
		url, err := mockUrlService.GetUrl("https://example.com/short")
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}

		if url != "https://example.com" {
			t.Errorf("Expected url to be %s, but got %s", "https://example.com", url)
		}
	})
	t.Run("Url not found", func(t *testing.T) {
		_, err := mockUrlService.GetUrl("https://example.com/short2")
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
		want := "Short URL not found"
		got := err.Error()
		if err.Error() != want {
			t.Errorf("Expected error to be %s, but got %s", want, got)
		}
	})

}

func TestUrlService_DeleteUrl(t *testing.T) {
	mockUrlService := &UrlServiceMock{
		DB: []*Url{
			{
				Url:      "https://example.com",
				ShortUrl: "https://example.com/short",
			},
		},
	}
	t.Run("Happy path", func(t *testing.T) {
		err := mockUrlService.DeleteUrl("https://example.com/short")
		if err != nil {
			t.Errorf("Expected no error, but got %s", err)
		}

		if len(mockUrlService.DB) != 0 {
			t.Errorf("Expected DB length to be 0, but got %d", len(mockUrlService.DB))
		}
	})
	t.Run("Url not found", func(t *testing.T) {
		err := mockUrlService.DeleteUrl("https://example.com/short2")
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
		want := "Short URL not found"
		got := err.Error()
		if err.Error() != want {
			t.Errorf("Expected error to be %s, but got %s", want, got)
		}
	})
}

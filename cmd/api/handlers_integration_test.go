package main

import (
	"echo_url_shortner/data"
	"echo_url_shortner/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppIntegration_GetUrl(t *testing.T) {
	e := echo.New()

	cfg := data.DefaultTestPostgresConfig()
	db, err := data.Open(cfg)
	if err != nil {
		t.Errorf("Expected no error, but got %s", err)
	}
	urlService := models.UrlService{DB: db}

	app := App{
		UrlModel: &urlService,
	}

	t.Run("Happy path", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/abc123", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetParamNames("url")
		c.SetParamValues("abcdef")

		if assert.NoError(t, app.GetUrl(c)) {
			assert.Equal(t, http.StatusMovedPermanently, rec.Code)
		}
	})
	t.Run("Get URL Not Found", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/abc123", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if err := app.GetUrl(c); err != nil {
			httpError, ok := err.(*echo.HTTPError)
			if ok {
				// Assert the HTTP status code
				assert.Equal(t, http.StatusNotFound, httpError.Code)

				// Assert the HTTP response body
				assert.Equal(t, "Url not found", httpError.Message)
			}
		}
	})
}

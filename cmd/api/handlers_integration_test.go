package main

import (
	"bytes"
	"echo_url_shortner/data"
	"echo_url_shortner/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApp_InsertUrl(t *testing.T) {
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
		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(map[string]string{
			"url": "https://testme.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", payload)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, app.InsertUrl(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var url models.Url
			if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&url)) {
				assert.Equal(t, "https://testme.com", url.Url)
			}
			err = urlService.DeleteUrl(url.ShortUrl)
			if err != nil {
				t.Errorf("Expected no error, but got %s", err)
			}
		}
	})
	t.Run("Duplicate url", func(t *testing.T) {
		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(map[string]string{
			"url": "https://www.google.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", payload)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		_ = app.InsertUrl(c)
		if assert.NoError(t, app.InsertUrl(c)) {
			assert.Equal(t, http.StatusConflict, rec.Code)
		}
	})

}

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

func TestAppIntegration_GetRandomLeetCode(t *testing.T) {
	cfg := data.DefaultTestPostgresConfig()
	db, err := data.Open(cfg)
	if err != nil {
		t.Error("unable to connect to postgres")
	}
	urlService := &models.UrlService{DB: db}
	app := App{UrlModel: urlService}
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/leetcode", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	assert.NoError(t, app.GetRandomLeetCode(c))
	assert.Equal(t, rec.Code, http.StatusFound)
}

package main

import (
	"bytes"
	"echo_url_shortner/models"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestInsertUrl(t *testing.T) {
	e := echo.New()

	mockUrlService := &models.UrlServiceMock{
		DB: []*models.Url{},
	}

	app := App{
		UrlModel: mockUrlService,
	}
	t.Run("Happy path", func(t *testing.T) {
		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(map[string]string{
			"url": "https://example.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", payload)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, app.InsertUrl(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var url models.Url
			if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&url)) {
				assert.Equal(t, "https://example.com", url.Url)
			}

			assert.Equal(t, 1, len(mockUrlService.DB))
			assert.Equal(t, "https://example.com", mockUrlService.DB[0].Url)
		}
	})

	t.Run("Bad request", func(t *testing.T) {

		// Send a JSON with the "urlx" field instead of "url"
		payload := bytes.NewBuffer([]byte(`{"urlx": "https://example.com"}`))
		req := httptest.NewRequest(http.MethodPost, "/", payload)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		_ = app.InsertUrl(c)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		// Assert the HTTP response body
		assert.Equal(t, "Bad request", rec.Body.String())
	})
	t.Run("Duplicate URL", func(t *testing.T) {
		// i need to set up a new mock user service and app here, there's still hold overs from the previous test
		mockUrlService := &models.UrlServiceMock{
			DB: []*models.Url{},
		}

		app := App{
			UrlModel: mockUrlService,
		} // First insert
		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(map[string]string{
			"url": "https://example.com",
		})
		req := httptest.NewRequest(http.MethodPost, "/", payload)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		if assert.NoError(t, app.InsertUrl(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

			var url models.Url
			if assert.NoError(t, json.NewDecoder(rec.Body).Decode(&url)) {
				assert.Equal(t, "https://example.com", url.Url)
			}

			assert.Equal(t, 1, len(mockUrlService.DB))
			assert.Equal(t, "https://example.com", mockUrlService.DB[0].Url)
		}

		// Second insert
		payload2 := new(bytes.Buffer)
		json.NewEncoder(payload2).Encode(map[string]string{
			"url": "https://example.com",
		})
		req2 := httptest.NewRequest(http.MethodPost, "/", payload2)
		req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec2 := httptest.NewRecorder()

		c2 := e.NewContext(req2, rec2)

		_ = app.InsertUrl(c2)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusConflict, rec2.Code)

		// Assert the HTTP response body
		assert.Equal(t, "Record already exists", rec2.Body.String())
	})
}

func TestApp_GetUrl(t *testing.T) {
	e := echo.New()

	mockUrlService := &models.UrlServiceMock{
		DB: []*models.Url{
			{
				Url:       "https://example.com",
				ShortUrl:  "abc123",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	app := App{
		UrlModel: mockUrlService,
	}
	t.Run("Happy path", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/abc123", nil)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		c.SetParamNames("url")
		c.SetParamValues("abc123")

		if assert.NoError(t, app.GetUrl(c)) {
			assert.Equal(t, http.StatusMovedPermanently, rec.Code)
		}
	})
	t.Run("Get URL Not Found", func(t *testing.T) {
		e := echo.New()

		mockUrlService := &models.UrlServiceMock{
			DB: []*models.Url{},
		}

		app := App{
			UrlModel: mockUrlService,
		}

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

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRoutes(t *testing.T) {
	e := echo.New()

	app := App{}

	app.SetupRoutes(e)

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(t, helloWorldHandler(c)) {
		// Check the status code and body
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello, World!", rec.Body.String())
	}
}

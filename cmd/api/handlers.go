package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func notFoundHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "The content you are looking for does not exist")
}

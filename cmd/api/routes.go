package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupRoutes(e *echo.Echo) {
	e.GET("/", helloWorldHandler)
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == http.StatusNotFound {
				notFoundHandler(c)
			} else {
				e.DefaultHTTPErrorHandler(err, c)
			}
		}
	}
}

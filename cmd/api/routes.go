package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (app App) SetupRoutes(e *echo.Echo) {
	e.GET("/", helloWorldHandler)
	e.POST("/urls", app.InsertUrl)
	e.GET("/:url", app.GetUrl)
	e.GET("/leetcode", app.GetRandomLeetCode)
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

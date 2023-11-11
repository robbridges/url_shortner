package main

import (
	"echo_url_shortner/models"
	"errors"
	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func helloWorldHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func notFoundHandler(c echo.Context) error {
	return c.String(http.StatusNotFound, "The content you are looking for does not exist")
}

func (app App) InsertUrl(c echo.Context) error {
	var payload struct {
		Url string `json:"url"`
	}

	err := c.Bind(&payload)
	if err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	url := &models.Url{
		Url:       payload.Url,
		ShortUrl:  StringWithCharset(6, charset),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = app.UrlModel.InsertUrl(url)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.ConstraintName == "url_url_key" {
				return c.String(http.StatusConflict, "Record already exists")
			}
		}
		return c.String(http.StatusInternalServerError, "error inserting url")
	}

	return c.JSON(http.StatusOK, url)
}

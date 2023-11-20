package main

import (
	"echo_url_shortner/data"
	"echo_url_shortner/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log"
	"log/slog"
)

type App struct {
	UrlModel models.IUrlService
}

func main() {
	viper.SetConfigFile("local.env")
	viper.AddConfigPath("/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("init: %w", err))
	}

	e := echo.New()

	cfg := data.DefaultPostgresConfig()
	db, err := data.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	app := App{}
	app.UrlModel = &models.UrlService{DB: db}

	app.SetupRoutes(e)
	log.Printf("Starting server at port 8080")
	err = e.Start(":8080")
	if err != nil {
		slog.Error("Error starting server")
	}

}

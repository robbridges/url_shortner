package main

import (
	"database/sql"
	"echo_url_shortner/cmd/api"
	"echo_url_shortner/data"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"log/slog"
)

type App struct {
	DB *sql.DB
}

func init() {
	viper.SetConfigFile("local.env")
	viper.AddConfigPath("/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("init: %w", err))
	}
}

func main() {
	e := echo.New()

	cfg := data.DefaultPostgresConfig()
	db, err := data.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	api.SetupRoutes(e)

	err = e.Start(":8080")
	if err != nil {
		slog.Error("Error starting server")
	}

}

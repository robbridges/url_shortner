package data

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
	"os"
)

type PostgressConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMODE  string
}

func DefaultPostgresConfig() PostgressConfig {
	return PostgressConfig{
		Host:     viper.GetString("DB_HOST"),
		Port:     viper.GetString("DB_PORT"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
		Database: viper.GetString("DATABASE"),
	}
}

func Open(config PostgressConfig) (*sql.DB, error) {

	db, err := sql.Open(
		"pgx",
		config.String(),
	)

	if err != nil {
		return nil, fmt.Errorf("error Opening DB: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func OpenDSN(dsn string) (*sql.DB, error) {
	connectionString := os.Getenv(dsn)
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("error Opening DB: %w", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Failed to ping the database:", err)
	}

	return db, nil
}

func (cfg PostgressConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMODE)
}

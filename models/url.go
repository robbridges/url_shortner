package models

import (
	"database/sql"
	"github.com/jackc/pgconn"
	"time"
)

type IUrlService interface {
	InsertUrl(url *Url) error
	GetUrl(short_url string) (string, error)
}

type Url struct {
	ID        int
	Url       string
	ShortUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UrlService struct {
	DB *sql.DB
}

type UrlServiceMock struct {
	DB []*Url
}

func (us *UrlService) InsertUrl(url *Url) error {
	stmt := `INSERT INTO url (url, short_url, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, url, short_url`
	err := us.DB.QueryRow(stmt, url.Url, url.ShortUrl, time.Now(), time.Now()).Scan(&url.ID, &url.Url, &url.ShortUrl)
	if err != nil {
		return err
	}
	return nil
}

func (us *UrlService) GetUrl(short_url string) (string, error) {
	var url string
	stmt := `SELECT url FROM url WHERE short_url = $1`
	err := us.DB.QueryRow(stmt, short_url).Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (m *UrlServiceMock) InsertUrl(url *Url) error {
	for _, u := range m.DB {
		if u.Url == url.Url {
			return &pgconn.PgError{ConstraintName: "url_url_key"}
		}
	}

	m.DB = append(m.DB, url)

	return nil
}

func (m *UrlServiceMock) GetUrl(short_url string) (string, error) {
	for _, url := range m.DB {
		if url.ShortUrl == short_url {
			return url.Url, nil
		}
	}
	return "", nil
}

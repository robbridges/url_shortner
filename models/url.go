package models

import (
	"database/sql"
	"fmt"
	"github.com/jackc/pgconn"
	"time"
)

type IUrlService interface {
	InsertUrl(url *Url) error
	GetUrl(short_url string) (string, error)
	DeleteUrl(short_url string) error
}

type Url struct {
	ID        int
	Url       string
	ShortUrl  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Leetcode  bool
}

type UrlService struct {
	DB *sql.DB
}

type UrlServiceMock struct {
	DB []*Url
}

func (us *UrlService) InsertUrl(url *Url) error {
	stmt := `INSERT INTO url (url, short_url, created_at, updated_at, leetcode) VALUES ($1, $2, $3, $4, $5) RETURNING id, url, short_url, leetcode`
	err := us.DB.QueryRow(stmt, url.Url, url.ShortUrl, time.Now(), time.Now(), url.Leetcode).Scan(&url.ID, &url.Url, &url.ShortUrl, &url.Leetcode)
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

func (m *UrlService) DeleteUrl(short_url string) error {
	statement := `DELETE FROM url WHERE short_url = $1`
	res, err := m.DB.Exec(statement, short_url)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("URL not found")
	}

	return nil
}

func (m *UrlServiceMock) GetUrl(short_url string) (string, error) {
	for _, url := range m.DB {
		if url.ShortUrl == short_url {
			return url.Url, nil
		}
	}
	return "", fmt.Errorf("Short URL not found")
}

func (m *UrlServiceMock) DeleteUrl(short_url string) error {
	for i, url := range m.DB {
		if url.ShortUrl == short_url {
			m.DB = append(m.DB[:i], m.DB[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Short URL not found")
}

package queries

import (
	"fmt"
	"time"

	db "github.com/harshitsinghai/felix/database"
	"github.com/harshitsinghai/felix/models"
)

// InsertURL into the database
func InsertURL(url *models.Record) (int, error) {
	var id int
	query := "INSERT INTO felix_url_shortner(original_url, short_url, created_at, expires_at, has_expired) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	insertError := db.DB.QueryRow(query, url.OriginalURL, url.ShortURL, url.CreatedAt, url.ExpiresAt, url.HasExpired).Scan(&id)
	return id, insertError
}

// FetchOriginalURL fetches the original URL from the given shortURL
func FetchOriginalURL(shortURL string) (bool, string) {
	var originalURL string
	var expireDate time.Time

	err := db.DB.QueryRow("SELECT original_url, expires_at FROM felix_url_shortner WHERE short_url = $1", shortURL).Scan(&originalURL, &expireDate)
	if err != nil {
		return false, "no rows found"
	}

	loc, err := time.LoadLocation("Asia/Kolkata")

	nowTime := time.Now().In(loc)

	hasExpired := expireDate.After(nowTime)
	has2Expired := nowTime.After(expireDate)
	fmt.Println("LINK has expired ", hasExpired, has2Expired)

	return true, originalURL
}

// FetchAlreadyExists Check if timestamp already exists
func FetchAlreadyExists(shortURL string) (bool, string) {
	var originalURL string
	var expireDate time.Time

	err := db.DB.QueryRow("SELECT original_url, expires_at FROM felix_url_shortner WHERE short_url = $1", shortURL).Scan(&originalURL, &expireDate)
	if err != nil {
		return false, "no rows found"
	}
	return true, originalURL
}

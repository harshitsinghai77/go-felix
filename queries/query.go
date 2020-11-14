package queries

import (
	db "github.com/harshitsinghai/felix/database"
	"github.com/harshitsinghai/felix/models"
)

// InsertURL into the database
func InsertURL(url *models.Record) (int, error) {

	var id int
	query := "INSERT INTO web_url(original_url, short_url, created_at, expires_at, has_expired) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	insertError := db.DB.QueryRow(query, url.OriginalURL, url.ShortURL, url.CreatedAt, url.ExpiresAt, url.HasExpired).Scan(&id)
	return id, insertError

}

// FetchOriginalURL fetches the original URL from the given shortURL
func FetchOriginalURL(shortURL string) (bool, string, error) {

	query := `SELECT original_url FROM web_url WHERE short_url = '` + shortURL + `' LIMIT 1;`
	rows, queryError := db.DB.Query(query)

	if queryError != nil {
		return false, "", queryError
	}
	defer rows.Close()

	for rows.Next() {
		originalURL := ""
		err := rows.Scan(&originalURL)

		if err != nil {
			return false, "", err
		}
		return true, originalURL, nil

	}
	return false, "No Url Found", nil

}

// FetchShortURLExists check if the given string already exists in the database
func FetchShortURLExists(shortURL string) (bool, string, error) {
	query := `SELECT original_url FROM web_url WHERE short_url = '` + shortURL + `' LIMIT 1;`

	rows, queryError := db.DB.Query(query)

	if queryError != nil {
		return false, "", queryError
	}

	defer rows.Close()

	for rows.Next() {
		originalURL := ""
		err := rows.Scan(&originalURL)

		if err != nil {
			return false, "", err
		}
		return true, originalURL, nil
	}

	return false, "", nil

}

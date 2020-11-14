package models

import "time"

// Record model contains the field expected by the database.
type Record struct {
	ID            int       `json:"id" db:"id"`
	OriginalURL   string    `json:"originalUrl" db:"original_url"`
	ShortURL      string    `json:"shortUrl" db:"short_url"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	ExpiresAt     time.Time `json:"expiresAfter" db:"expires_at"`
	HasExpired    bool      `json:"hasExpired" db:"has_expired"`
	AlreadyExists bool      `json:"alreadyExists"`
}

// GenerateRequest is the expected POST json request from the user
type GenerateRequest struct {
	OriginalURL  string `json:"originalUrl"`
	ExpiresAfter string `json:"expiresAfter"`
}

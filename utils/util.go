package utils

import (
	"time"
)

// SetExpiryDate sets the expiry date based on given hour minute or second
func SetExpiryDate(expiryType string, duration uint64) time.Time {

	// loc, _ := time.LoadLocation("Asia/Kolkata")

	switch expiryType {

	case "day":
		return time.Now().Add(time.Hour * 24 * time.Duration(duration))
	case "hour":
		return time.Now().Add(time.Hour * time.Duration(duration))
	case "min":
		return time.Now().Add(time.Minute * time.Duration(duration))
	}

	return time.Now().UTC()
}

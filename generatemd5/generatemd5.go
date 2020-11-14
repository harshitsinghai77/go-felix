package generatemd5

import (
	"crypto/md5"
	"encoding/hex"
)

// GenerateHash generates hash from the string
func GenerateHash(originalURL string) string {
	md5Sum := md5.Sum([]byte(originalURL))
	md5Hash := hex.EncodeToString(md5Sum[:])
	shortURL := md5Hash[:6]
	return shortURL
}

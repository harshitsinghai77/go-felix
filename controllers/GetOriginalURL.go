package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	query "github.com/harshitsinghai/felix/queries"
)

// GetOriginalURL fetches the original URL for the given encoded(short)
func GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var originalURL string
	vars := mux.Vars(r)

	shortURL := vars["shortUrl"]

	// Get ID from base62 string
	// id := generatemd5.GenerateShortURL(vars["encoded_url"])
	exists, originalURL, err := query.FetchOriginalURL(shortURL)

	// Handle response details
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responseMap := map[string]interface{}{"status": exists, "originalURL": originalURL}
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

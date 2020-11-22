package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/harshitsinghai/felix/models"
	query "github.com/harshitsinghai/felix/queries"
)

// GetStyleTransferImages fetches the original style transfer image URL from the database
func GetStyleTransferImages(w http.ResponseWriter, r *http.Request) {

	exists, allImages := query.FetchStyleTransferURL()

	// Handle response details
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No data found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responseMap := map[string]interface{}{"status": exists, "data": allImages}
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

// InsertImageURL inserts image URL in the database
func InsertImageURL(w http.ResponseWriter, r *http.Request) {

	var requestBody models.ImageURL

	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &requestBody)

	id, err := query.InsertStyleTransferURL(requestBody)

	// Handle response details
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Some error occured"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	responseMap := map[string]interface{}{"status": true, "id": id}
	response, _ := json.Marshal(responseMap)
	w.Write(response)
}

package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	controller "github.com/harshitsinghai/felix/controllers"
	database "github.com/harshitsinghai/felix/database"
	"github.com/rs/cors"
)

// HelloInit makes sure the server is connected to the client
func HelloInit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!!"))
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	database.TestPing()
	if err != nil {
		panic(err)
	}

	defer db.Close(context.Background())

	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloInit).Methods("GET")
	r.HandleFunc("/short", controller.GenerateShortURL).Methods("POST")
	r.HandleFunc("/short/{shortUrl}", controller.GetOriginalURL).Methods("GET")

	corsHandler := cors.AllowAll().Handler(r)

	srv := &http.Server{
		Handler: corsHandler,
		Addr:    "localhost:8000",

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening to server 8000")
	log.Fatal(srv.ListenAndServe())
}

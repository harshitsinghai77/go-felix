package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	controller "github.com/harshitsinghai/felix/controllers"
	database "github.com/harshitsinghai/felix/database"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// HelloInit makes sure the server is connected to the client
func HelloInit(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Go!!"))
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	fmt.Println(PORT)
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	database.TestPing()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/hello", HelloInit).Methods("GET")
	r.HandleFunc("/short", controller.GenerateShortURL).Methods("POST")
	r.HandleFunc("/short/{shortUrl}", controller.GetOriginalURL).Methods("GET")

	corsHandler := cors.AllowAll().Handler(r)

	srv := &http.Server{
		Handler: corsHandler,
		Addr:    "0.0.0.0:" + PORT,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

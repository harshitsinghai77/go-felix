package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

// DB variable holds database property
var DB *pgx.Conn

// DBerr is the error of the database
var DBerr error

// InitDB initialized the database and creates a WEB_URL table
func InitDB() (*pgx.Conn, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	var err error

	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname := os.Getenv("PG_DATABASE")

	DatabaseURI := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname
	DB, DBerr = pgx.Connect(context.Background(), DatabaseURI)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	table, err := DB.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS web_url( 
		    id SERIAL,
		    original_url varchar not null,
			short_url varchar PRIMARY KEY not null,
			created_at time not null,
			expires_at time,
			has_expired bool default false);`)

	if err != nil {
		log.Println(table)
		log.Println(err)
	}
	DB.Exec(context.Background(), "create unique index shortUrl_original_index on web_url(original_url);")
	DB.Exec(context.Background(), "create unique index shortUrl_short_index on web_url(short_url);")

	return DB, nil

}

// TestPing test function for the package database
func TestPing() {

	// Pings the global database
	pingError := DB.Ping(context.Background())

	if pingError != nil {
		// An error was returned while pinging the database
		log.Fatal(pingError)
	} else {
		// Database Ping successful
		log.Println("Database: Ping successful.")
	}

}

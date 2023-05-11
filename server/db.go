package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func getDb() *sql.DB {
	connStr := "postgresql://postgres:postgres@localhost/postgres?sslmode=disable"
	// Connect to database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

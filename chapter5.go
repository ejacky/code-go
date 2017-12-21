package main

import (
	"database/sql"
	"github.com/lib/pg"
)

func main() {
	conStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("skylar", connStr)
	if err != nil {
		log.Fatal(err)
	}

	id := 142093
	rows, err := db.Query("SELECT name FROM client WHERE id = $1", id)
}
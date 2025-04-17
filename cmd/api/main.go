package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const connStr = "postgres://ticket_owner:npg_k52oIrYlRgwz@ep-cold-art-a4q6c276-pooler.us-east-1.aws.neon.tech/ticket?sslmode=require"

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	log.Println("Connected to PostgreSQL database successfully!")
}

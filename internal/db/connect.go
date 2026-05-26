package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ncruces/go-sqlite3/driver"
)

func ConnectDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "database/database.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("could not connect to sqlite: %v", err)
	}

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	fmt.Println("success connect database :)")
	return db, nil
}

func InitSchema(db *sql.DB) {
	schema, err := os.ReadFile("internal/db/schema.sql")
	if err != nil {
		log.Fatalf("read schema failed: %v", err)
	}

	if _, err := db.Exec(string(schema)); err != nil {
		log.Fatalf("apply schema failed: %v", err)
	}
}

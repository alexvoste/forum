package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/internal/db"
	"forum/internal/routes"
)

func main() {
	database, err := db.ConnectDatabase()
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	defer database.Close()
	db.InitSchema(database)

	fmt.Println("server starting on https://localhost:8080 lol")
	routes.InitRoutes(database)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server failed to start: %v\n", err)
	}
}

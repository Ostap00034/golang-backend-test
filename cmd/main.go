package main

import (
	"database/sql"
	"log"
	"main/cmd/api"
	"main/config"
	"main/db"
)

func main() {
	connStr := config.ConnStr

	db, err := db.NewPostgreSQLStorage(connStr)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected")
}

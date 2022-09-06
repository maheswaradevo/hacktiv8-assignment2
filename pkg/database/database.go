package database

import (
	"database/sql"
	"fmt"
	"log"
)

func GetDatabase(dbAddress string, dbUsername string, dbPassword string, dbName string) *sql.DB {
	log.Printf("INFO GetDatabase database connection: starting database connection process")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		dbUsername, dbPassword, dbAddress, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatalf("Error GetDatabase sql open connection fatal error: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("ERROR GetDatabase db ping fatal error: %v", err)
	}
	log.Printf("INFO GetDatabase database connectionn: established successfully with %s\n", dataSourceName)
	return db
}

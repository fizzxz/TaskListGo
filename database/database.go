package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func Close(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		return err
	}
	return nil
}

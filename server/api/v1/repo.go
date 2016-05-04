package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Initialise Database
func init() {
	connectToDatabse()
}

//connectToDatabse does what you expect it to do. Set enviroment variables first
func connectToDatabse() {

	// connect to the database
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

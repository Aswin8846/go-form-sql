package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	// Connect to MySQL
	DB, err = sql.Open("mysql", "root:haywin$8846@tcp(localhost:3306)/")
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Create database if not exists
	_, err = DB.Exec("CREATE DATABASE IF NOT EXISTS testdb")
	if err != nil {
		log.Fatal("Failed to create database:", err)
	}

	// Use the testdb database
	_, err = DB.Exec("USE testdb")
	if err != nil {
		log.Fatal("Failed to select database:", err)
	}

	// Reconnect pointing to the database
	DB, err = sql.Open("mysql", "root:haywin$8846@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatal("Failed to connect to testdb:", err)
	}

	// Test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	log.Println("Database connected successfully!")
}


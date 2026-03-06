package main

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	dbPath := filepath.Join(os.Getenv("DB_PATH"), "tasks.db")
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping DB:", err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		status TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
	log.Println("Database initialized successfully!")
	return db
}

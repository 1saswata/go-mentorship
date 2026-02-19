package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "./tasks.db")
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

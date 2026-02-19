package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func InitDB() {
	db, err := sql.Open("sqlite", "./tasks.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB connected!")
	}
	res, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		status TEXT NOT NULL
	);`)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}

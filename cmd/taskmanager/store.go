package main

import (
	"database/sql"
	"log"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TaskStore struct {
	db *sql.DB
}

func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{db: db}
}

func (t *TaskStore) CreateTask(name, status string) int {
	query := "INSERT INTO tasks (name, status) VALUES (?, ?)"
	result, err := t.db.Exec(query, name, status)
	if err != nil {
		log.Println("Error inserting task:", err)
		return -1
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting id:", err)
		return -1
	}
	return int(id)
}

func (t *TaskStore) GetAllTasks() []Task {
	rows, err := t.db.Query("Select id, name, status from tasks")
	if err != nil {
		log.Println("Error getting tasks:", err)
		return nil
	}
	defer rows.Close()
	allTasks := []Task{}
	for rows.Next() {
		var t Task
		if err = rows.Scan(&t.ID, &t.Name, &t.Status); err != nil {
			log.Println("Error scanning row: ", err)
			continue
		}
		allTasks = append(allTasks, t)
	}
	return allTasks
}

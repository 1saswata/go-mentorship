package main

import "sync"

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type TaskStore struct {
	sync.Mutex
	tasks  map[int]Task
	nextID int
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		tasks:  make(map[int]Task),
		nextID: 1,
	}
}

func (t *TaskStore) CreateTask(name, status string) int {
	t.Lock()
	defer t.Unlock()
	task := Task{
		ID:     t.nextID,
		Name:   name,
		Status: status,
	}
	t.tasks[t.nextID] = task
	t.nextID++
	return task.ID
}

func (t *TaskStore) GetAllTasks() []Task {
	t.Lock()
	defer t.Unlock()
	allTasks := make([]Task, 0, len(t.tasks))
	for _, task := range t.tasks {
		allTasks = append(allTasks, task)
	}
	return allTasks
}

package models

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Done      bool      `json:"done"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func done(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

func Tasks() []Task {
	db := Connect()
	// rows, err := db.Query("SELECT id, state, body, created_at, updated_at FROM tasks WHERE state = $1", 0)
	rows, err := db.Query("SELECT id, state, body, created_at, updated_at FROM tasks")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var id, state int
		var body string
		var created_at, updated_at time.Time
		if err := rows.Scan(&id, &state, &body, &created_at, &updated_at); err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		task := Task{ID: id, Done: done(state), Body: body, CreatedAt: created_at, UpdatedAt: updated_at}
		fmt.Printf("task: %v", task)
		tasks = append(tasks, task)
	}

	return tasks
}

func FindTask(id string) Task {
	var task Task
	db := Connect()
	rows, err := db.Query("SELECT id, state, body, created_at, updated_at FROM tasks WHERE id = $1", id)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer rows.Close()

	if rows.Next() {
		var id, state int
		var body string
		var created_at, updated_at time.Time
		if err := rows.Scan(&id, &state, &body, &created_at, &updated_at); err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		task = Task{ID: id, Done: done(state), Body: body, CreatedAt: created_at, UpdatedAt: updated_at}
	} else {
		panic(fmt.Sprintf("Task should return a valid entry"))
	}

	return task
}

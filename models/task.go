package models

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Task struct {
	ID        int64     `json:"id"`
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
	db := DB()
	stmt, err := db.Prepare("SELECT id, state, body, created_at, updated_at FROM tasks")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer rows.Close()

	tasks := make([]Task, 0)
	for rows.Next() {
		var id int64
		var state int
		var body string
		var created_at, updated_at time.Time
		if err := rows.Scan(&id, &state, &body, &created_at, &updated_at); err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		task := Task{ID: id, Done: done(state), Body: body, CreatedAt: created_at, UpdatedAt: updated_at}
		fmt.Printf("task: %v", task)
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return tasks
}

func FindTask(id string) Task {
	var task Task
	db := DB()
	stmt, err := db.Prepare("SELECT id, state, body, created_at, updated_at FROM tasks WHERE id = $1")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	defer rows.Close()

	if rows.Next() {
		var id int64
		var state int
		var body string
		var created_at, updated_at time.Time
		if err := rows.Scan(&id, &state, &body, &created_at, &updated_at); err != nil {
			panic(fmt.Sprintf("%v", err))
		}
		task = Task{ID: id, Done: done(state), Body: body, CreatedAt: created_at, UpdatedAt: updated_at}
	} else {
		panic(fmt.Sprintf("Task should return a valid entry"))
	}

	err = rows.Err()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return task
}

func CreateTask(input io.Reader) Task {
	type jsonFmt struct {
		Task map[string]string
	}
	var lastId int64

	//get content of JSON element
	decoder := json.NewDecoder(input)
	var task jsonFmt
	err := decoder.Decode(&task)
	if err != nil {
		panic("Unable to decode: " + err.Error())
	}
	body := task.Task["body"]

	// ADD to database
	db := DB()
	stmt, err := db.Prepare("INSERT INTO tasks(body, state, created_at, updated_at) VALUES($1, $2, $3, $4) RETURNING id")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	err = stmt.QueryRow(body, 0, time.Now(), time.Now()).Scan(&lastId)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	// get database details
	rowCnt, err := res.RowsAffected()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	fmt.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	// return the created task
	return Task{ID: lastId, Done: false, Body: body}
}

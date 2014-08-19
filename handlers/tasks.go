package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	// this project
	"github.com/itchy/go-ng-todo/models"
)

func TasksHandler(w http.ResponseWriter, req *http.Request, params map[string]string) {
	// GET data
	tasks := models.Tasks()
	b, _ := json.Marshal(tasks)
	// Write output to http.ResponseWriter
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
	return
}

func CreateTaskHandler(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Println("\n\nCreateTaskHandler\n\n")

	task := Response{"id": "1", "body": "sample task"}
	w.Header().Set("Content-Type", "application/json")
	// fmt.Fprint(w, Response{"success": true, "message": "Hello!"})
	fmt.Fprint(w, task)
	return
}

func UpdateTaskHandler(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Println("\n\nUpdateTaskHandler\n\n")
	task_id := string(params["id"][0])
	task := models.FindTask(task_id)

	fmt.Println("\n\n %v \n\n", task)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "")
	return
}

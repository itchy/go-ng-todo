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
	// PASS TO MODEL
	_ = models.CreateTask(req.Body)

	// return list of all tasks
	tasks := models.Tasks()
	b, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
	return
}

func UpdateTaskHandler(w http.ResponseWriter, req *http.Request, params map[string]string) {
	fmt.Println("\n\nUpdateTaskHandler\n\n")
	task_id := string(params["id"][0]) // just grabbing first element before .json need a better way to do this

	task := models.FindTask(task_id)
	b, _ := json.Marshal(task)

	fmt.Println("Query String")
	fmt.Println(req.URL.RawQuery)
	fmt.Println("Element JSON")
	fmt.Println(string(b))
	fmt.Println("")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
	return
}

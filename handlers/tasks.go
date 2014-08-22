package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	// this project
	"github.com/itchy/go-ng-todo/models"
)

func TasksHandler(w http.ResponseWriter, req *http.Request, params map[string]string) {
	log("INFO", "Processing tasks handler")
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
	task_id := string(strings.Split(params["id"], ".")[0]) // strip off .json
	fmt.Println(task_id)

	err := models.UpdateTask(task_id, req.Body)
	if err != nil {
		panic("Unable to update task: " + err.Error())
	}

	task := models.FindTask(task_id)
	b, _ := json.Marshal(task)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
	return
}

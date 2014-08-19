package config

/*
	The rouets file is a part of the config
	it is were all routes are passed to handlers
	for the application
*/

import (
	"github.com/dimfeld/httptreemux"
	// this project
	"github.com/itchy/go-ng-todo/handlers"
)

func Router() *httptreemux.TreeMux {
	router := httptreemux.New()

	// set up routes
	// router.GET("/", handlers.HomeHandler)
	router.GET("/tasks.json", handlers.TasksHandler)
	router.POST("/tasks.json", handlers.CreateTaskHandler)
	router.POST("/tasks/:id", handlers.UpdateTaskHandler)
	router.PUT("/tasks/:id", handlers.UpdateTaskHandler)
	router.PATCH("/tasks/:id", handlers.UpdateTaskHandler)

	// pass back to config
	return router
}

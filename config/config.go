package config

/*
	This is where we configure /setup/ the app
	set things like
  - router
  - port
  - logger
  - middle ware
  - defaults
*/

import (
	"github.com/codegangsta/negroni"
	// this project
	"github.com/itchy/go-ng-todo/database"
)

func Run(configs map[string]string) {
	router := Router()
	database.Connect()
	defer database.Close()
	// logger
	logger := Logger(configs["logFile"])

	// config negroni
	n := negroni.Classic()
	n.UseHandler(router)
	n.Use(logger)
	n.Run(":" + configs["port"])
}

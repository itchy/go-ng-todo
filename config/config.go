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
	"github.com/itchy/go-ng-todo/utils"
)

func Run(configs map[string]string) {
	router := Router()
	database.Connect()
	defer database.Close()
	// logger
	utils.SetLogger(configs["logFile"])

	// config negroni
	n := negroni.Classic()
	n.UseHandler(router)
	n.Use(utils.Logger())
	n.Run(":" + configs["port"])
}

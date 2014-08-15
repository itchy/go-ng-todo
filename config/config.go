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
)

func Run(configs map[string]string) {
	router := Router()

	// config negroni
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":" + configs["port"])
}

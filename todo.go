package main

import (
	// This project
	"github.com/itchy/go-ng-todo/config"
)

func main() {
	configutation := map[string]string{
		"port": "3000",
		"data": "/var/todo/data"}

	config.Run(configutation)
}

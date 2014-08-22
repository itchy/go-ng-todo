package utils

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"log"
	"os"
)

var logger *negroni.Logger

func Logger() *negroni.Logger {
	return logger
}

func SetLogger(fileName string) {
	f, err := os.OpenFile("log/"+fileName+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: %v", err)
	}
	logger = negroni.NewLogger()
	logger.Logger = log.New(f, "[todo]", log.LstdFlags|log.Lshortfile)

	return
}

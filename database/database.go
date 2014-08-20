package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// TODO Need to make this long living!
// @db ||= sql.Open(...)

var db *sql.DB

func Connect() {
	fmt.Println("\n\nCONNECTING TO DATABASE\n\n")
	var err error
	db, err = sql.Open("postgres", "user=scott dbname=todo sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
}

func DB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}

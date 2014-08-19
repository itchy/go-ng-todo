package models

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// TODO Need to make this long living!
// @db ||= sql.Open(...)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "user=scott dbname=todo sslmode=disable")
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	return db
}

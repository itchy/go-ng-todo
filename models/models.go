package models

import (
	"database/sql"
	// this project
	"github.com/itchy/go-ng-todo/database"
)

func DB() *sql.DB {
	return database.DB()
}

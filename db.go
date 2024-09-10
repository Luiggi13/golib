package golib

import (
	"database/sql"
	"fmt"
)

func connect_db(connection string) *sql.DB {
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("failed connection!")
		panic(err)
	}

	return db
}

type LogInsert struct {
	Details string
	Path    string
	Method  string
}

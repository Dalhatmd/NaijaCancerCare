package database

import (
	"log"
	"database/sql"

	_"github.com/lib/pq"
)

func Connect(dbUrl string) *sql.DB {

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

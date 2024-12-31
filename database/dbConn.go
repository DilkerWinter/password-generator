package databaseConn

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() (*sql.DB, error) {
	connStr := "postgres://postgres:123@localhost:5432/passwordgenerator?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
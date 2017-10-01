package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:secret@localhost/rest?sslmode=disable")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to database")

	return db, nil
}

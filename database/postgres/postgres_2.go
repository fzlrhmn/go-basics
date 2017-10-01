package main

import (
	"fmt"
	"log"

	"github.com/fzlrhmn/go-basics/database/postgres/connection"
)

type User struct {
	id       string
	username string
	fullname string
	address  string
}

func main() {
	db, err := connection.NewConn("postgres://postgres:secret@localhost/rest?sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, username, fullname, address FROM users;")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.id, &user.username, &user.fullname, &user.address) // order matters
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("%s, %s, %s, %s\n", user.id, user.username, user.fullname, user.address)
	}
}

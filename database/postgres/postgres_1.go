package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	id       string
	username string
	fullname string
	address  string
}

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:secret@localhost/rest?sslmode=disable")
	if err != nil {
		log.Fatal("Error: The data source arguments are not valid")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error: Couldn't connect to database")
	}

	fmt.Println("Connected to database")

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

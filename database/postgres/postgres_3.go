// Using query builder to get data instead of write manual query

package main

import (
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/fzlrhmn/go-basics/database/postgres/connection"
)

type User struct {
	id       string
	username string
	fullname string
	address  string
}

func main() {
	db, err := connection.NewConn()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := sq.Select("id, username, fullname, address").From("users").RunWith(db).Query()
	if err != nil {
		panic(err)
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

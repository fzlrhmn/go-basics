package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", createUser)

	http.ListenAndServe(":8082", mux)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user := User{}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	log.Println(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

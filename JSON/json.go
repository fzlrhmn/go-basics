package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	// create json
	m := Message{"faisal", "hello world", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Println(string(b))

	// parse json
	str := `{"Name":"faisal","Body":"hello world","Time":1294706395881547000}`
	res := Message{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.Name)
}

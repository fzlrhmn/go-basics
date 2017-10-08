package main

import (
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//bPath := filepath.Base("/data/data.txt")
	//checkErr(err)
	//fmt.Println(bPath)
	//dat, err := ioutil.ReadFile("data.txt")
	file, err := os.Open("data.txt")
	checkErr(err)

	defer file.Close()

	b1 := make([]byte, 10)
	n1, err := file.Read(b1)
	checkErr(err)

	fmt.Println("%d bytes: %s\n", n1, string(b1))
}

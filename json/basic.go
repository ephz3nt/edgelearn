package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// an instance of our Book struct
	book := Book{
		Title:  "Learning Concurrency in Python",
		Author: "Elliot Forbes",
	}
	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))
}

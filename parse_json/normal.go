package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

var jsonString = []byte(`
[
	{"title": "title 1", "author":"taro"},
	{"title": "title 2", "author":"jiro"}
]`)

func main() {
	var books []Book
	err := json.Unmarshal(jsonString, &books)
	if err != nil {
		panic(err)
	}

	for _, book := range books {
		fmt.Println(book)
	}
}

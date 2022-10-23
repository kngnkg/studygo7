package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	log.SetFlags(log.Llongfile)

	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			{
				Id:      3,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			{
				Id:      4,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create("13json_creating_encoder/post.json")
	if err != nil {
		log.Print("Error creating JSON file:", err)
		return
	}
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	err = encoder.Encode(&post)
	if err != nil {
		log.Print("Error encoding JSON to file:", err)
		return
	}
}

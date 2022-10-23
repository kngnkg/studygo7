package main

import (
	"encoding/xml"
	"log"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	log.SetFlags(log.Llongfile)

	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		log.Print(err)
		return
	}
	err = os.WriteFile("post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		log.Print("Error writing XML to file:", err)
		return
	}
}

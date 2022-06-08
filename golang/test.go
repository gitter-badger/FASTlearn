package main

import (
	"fmt"
	"log"
	"os"
)

type SongInfo []struct {
	Author     string   `json:"author"`
	Paragraphs []string `json:"paragraphs"`
	Tags       []string `json:"tags"`
	Title      string   `json:"title"`
	ID         string   `json:"id"`
}

func main() {
	InputFile, err := os.Open("")
	defer InputFile.Close()
	if err != nil {
		log.Fatal(err)
		fmt.Println("ERROR 001")
	} else {
		fmt.Println("Opened")
	}
}

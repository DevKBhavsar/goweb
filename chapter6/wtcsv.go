package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func writeCSV() {
	csvFile, err := os.Create("post.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)

	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err = writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

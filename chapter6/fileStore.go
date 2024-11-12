package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func readFromFile() {
	data := []byte("HELLO WORLD")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, err := ioutil.ReadFile("data1")
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(read1))

	file1, _ := os.Create("data2")
	defer file1.Close()

	bytes, _ := file1.Write(data)
	fmt.Printf("WRITE %d in file 1", bytes)

	file2, _ := os.Open("data2")
	defer file2.Close()

	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("WRITE %d in file 2", bytes)

	fmt.Println(string(read2))
}

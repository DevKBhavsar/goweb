package main

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
)



func store(data interface{}, fileName string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fileName, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, fileName string) {
	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	post := Post{Id: 1, Content: "HELLO DEV", Author: "DEV"}
// 	store(post, "post1")
// 	var postRead Post
// 	load(&postRead, "post1")
// 	fmt.Println(postRead)
// }

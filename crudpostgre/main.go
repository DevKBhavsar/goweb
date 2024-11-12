package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// returns ALL THe post
func Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id,content,author from fb")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		post := Post{}
		rows.Scan(&post.Id, &post.Content, &post.Author)
		posts = append(posts, post)
	}
	rows.Close()
	return

}

// get post by id
func GetPost(id int) (post Post, err error) {
	err = Db.QueryRow("select id,content,author from fb where id=$1", id).Scan(&post.Id, &post.Content, &post.Author)
	return

}

// create post
func (p *Post) create() (err error) {
	statement := "insert into fb (content,author) values ($1,$2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.Id)
	return

}

// update post
func (p *Post) update() (err error) {
	_, err = Db.Exec("update fb set content=$1,author=$2 where id=$3", p.Content, p.Author, p.Id)
	return
}

// delete post
func (p *Post) delete() (err error) {
	_, err = Db.Exec("delete from fb where id=$1", p.Id)
	return
}

func main() {
	post1 := Post{Content: "Post 1 contenet", Author: "DEV"}
	fmt.Println(post1)

	post1.create()
	fmt.Println(post1)

	temp, _ := GetPost(post1.Id)
	fmt.Println(temp)

	temp.Content = "Updated Contenet"
	temp.Author = "Updated Author"
	temp.update()

	posts, _ := Posts()
	fmt.Println(posts)

	post1.delete()

}

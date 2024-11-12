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

// type Comment struct{
// 	Id int
// 	Content string
// 	Author string
// 	Post *Post
// }

var Db *sql.DB

// initializing database
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts() (posts []Post, err error) {
	rows, err := Db.Query("select id,content,author from posts")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			panic(err)
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select idcontent,author from posts where id=$1", id).Scan(&post.Id, &post.Author, &post.Content)
	return

}

func (post *Post) Create() (err error) {
	statement := "insert into posts (content,author) values ($1,$2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = $2,author =$3 where id=$1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from posts where id= $1", post.Id)
	fmt.Println("DELTE POST SUCCES")
	return
}

func trunc() (err error) {
	_, err = Db.Exec("TRUNCATE TABLE posts")
	return
}
func main() {

	trunc()
	post := Post{Content: "Hello Post", Author: "DEV BHAVSAR"}
	fmt.Println(post)
	post.Create()
	fmt.Println(post)

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Updating Contnent"
	readPost.Author = "Family"
	readPost.update()

	checkingUpdate := Post{Content: "before Update", Author: "After Update"}
	checkingUpdate.Create()
	checkingUpdate.Content = "Updated Contnent"
	checkingUpdate.Author = "updated Author"
	checkingUpdate.update()

	posts, _ := Posts()
	fmt.Println(posts)

}

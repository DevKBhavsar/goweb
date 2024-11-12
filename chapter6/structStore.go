package main



var byId map[int]*Post
var byAuthor map[string][]*Post

func storeData(post Post) {
	byId[post.Id] = &post
	byAuthor[post.Author] = append(byAuthor[post.Author], &post)
}

// Code for main
// byId = make(map[int]*Post)
// 	byAuthor = make(map[string][]*Post)

// 	p1 := Post{id: 1, contnet: "THIS IS DEV POST", author: "dev"}
// 	p2 := Post{id: 2, contnet: "THIS IS BHai's POST", author: "bhai"}
// 	p3 := Post{id: 3, contnet: "THIS IS DEV POST", author: "dev"}

// 	store(p1)
// 	store(p2)
// 	store(p3)

// 	fmt.Println(byId[1])

// 	for _,post:= range byAuthor["dev"]{
// 		fmt.Println(post)
// 	}

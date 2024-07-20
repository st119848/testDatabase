package main

import (
	"fmt"
	"testDatabase/data"
	"time"
)

func main() {
	// fmt.Println("Hello, world!")
	data.Init()
	// posts, _ := data.AllBlogPosts()
	// fmt.Println(posts)
	p := data.Post{"Hello", "World", time.Now()}
	fmt.Println(p)
	err := data.AddPost(p)
	fmt.Println(err)
	p.Save()

}

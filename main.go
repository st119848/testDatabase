package main

import (
	"fmt"
	"testDatabase/data"
)

func main() {
	fmt.Println("Hello, world!")
	data.Init()
	posts := data.AllBlogPosts()
	fmt.Println(posts)

}

package main

import (
	"fmt"
	"testDatabase/data"

	"github.com/gin-gonic/gin"
)

func main() {
	data.InitGorm()
	defer data.CloseGorm()

	r := gin.Default()

	r.POST("/person", createPerson)
}
func createPerson(c *gin.Context) {
	var person data.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		fmt.Print(person)
	}
}
func mainOld() {
	// fmt.Println("Hello, world!")
	data.Init()
	// posts, _ := data.AllBlogPosts()
	// fmt.Println(posts)
	// p := data.Post{"Hello", "World", time.Now()}
	// fmt.Println(p)
	// err := data.AddPost(p)
	// fmt.Println(err)
	// p.Save()
	// post := data.PostById(1)
	// fmt.Print(post)
	r := gin.Default()
	posts, _ := data.AllBlogPosts()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"posts":   posts,
		})

	})
	r.Run()

}

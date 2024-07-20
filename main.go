package main

import (
	"fmt"
	"strconv"
	"testDatabase/data"

	"github.com/gin-gonic/gin"
)

func main() {
	data.InitGorm()
	defer data.CloseGorm()

	r := gin.Default()

	r.POST("/person", createPerson)
	r.PUT("/person/:id", updatePerson)

	r.Run()
}

func updatePerson(c *gin.Context) {
	id := c.Param("id")
	var person data.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		fmt.Println("Error creating person")
	}
	fmt.Println(person)
	myUint, _ := strconv.ParseUint(id, 10, 64)
	data.UpdatePerson(uint(myUint), person)
	c.JSON(200, person)
}

func createPerson(c *gin.Context) {
	var person data.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		fmt.Println("Error creating person")
	}
	fmt.Println(person)
	data.CreatePerson(person)
	c.JSON(200, person)
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

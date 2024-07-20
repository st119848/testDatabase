package data

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func InitGorm() {
	// db, err := gorm.Open("sqlite3", "./gorm.db")
	db, err := gorm.Open("mysql", "root:root@tcp(localhost)/post?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	// Auto migrate the schema
	db.AutoMigrate(&Person{})

	// Create some records
	p1 := Person{FirstName: "John", LastName: "Doe"}
	p2 := Person{FirstName: "Jane", LastName: "Smith"}
	db.Create(&p1)
	db.Create(&p2)

	// Retrieve the first record from the database
	var p3 Person
	db.First(&p3)

	// Print the results
	fmt.Println(p1.FirstName) // Should print "John"
	fmt.Println(p2.LastName)  // Should print "Smith"
	fmt.Println(p3.LastName)  // Should print the last name of the first record, which is "Doe"

}

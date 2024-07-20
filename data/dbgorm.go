package data

import (
	"log"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
}

func InitGorm() {
	// db, err := gorm.Open("sqlite3", "./gorm.db")
	db, err := gorm.Open("mysql", "root:root@tcp(localhost)/post?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&Person{})

}
func CloseGorm() {
	db.Close()
}

func test() {
	// // Create some records
	// p1 := Person{FirstName: "John", LastName: "Doe", NickName: "A"}
	// p2 := Person{FirstName: "Jane", LastName: "Smith", NickName: "B"}
	// db.Create(&p1)
	// db.Create(&p2)

	// // Retrieve the first record from the database
	// var p3 Person
	// db.First(&p3)

	// // Print the results
	// // fmt.Println(p1.FirstName) // Should print "John"
	// // fmt.Println(p2.LastName)  // Should print "Smith"
	// // fmt.Println(p3.LastName)  // Should print the last name of the first record, which is "Doe"
	// /// ReadAll
	// var people []Person
	// db.Find(&people)

	// fmt.Println(people)
	// // // Read one
	// var person Person
	// db.Where("id =?", 1).First(&person)
	// fmt.Println(person)
	// // // Update
	// person.NickName = "OKOKOKOK"
	// person.LastName = "BBBBBB"
	// person.FirstName = "AAAAAA"
	// db.Save(&person)
	// // // Delete
	// db.Delete(&person)
	// fmt.Println("Deleted")
}

package data

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	NickName  string `json:"nickname"`
}

func InitGorm() {
	var err error
	// db, err = gorm.Open("sqlite3", "./gorm.db")
	db, err := gorm.Open("mysql", "root:root@tcp(localhost)/post?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Auto migrate the schema
	db.AutoMigrate(&Person{})
}

func CloseGorm() {
	if db != nil {
		db.Close()
	}
}

func CreatePerson(person Person) error {
	db, _ := gorm.Open("mysql", "root:root@tcp(localhost)/post?charset=utf8&parseTime=True&loc=Local")

	if db == nil {
		fmt.Println("Error DB connection")
	}
	if err := db.Create(&person).Error; err != nil {
		return err
	}
	return nil
}

package data

import (
	_ "database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() {
	db = sqlx.MustConnect("mysql", "root:root@tcp(localhost:3306)/post?parseTime=true")
}

type Post struct {
	Title   string
	Body    string
	Created time.Time
}

func AllBlogPosts() ([]Post, error) {
	posts := []Post{}
	err := db.Select(&posts, "SELECT title, body, created FROM posts")
	return posts, err
}

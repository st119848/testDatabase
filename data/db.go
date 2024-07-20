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

func AllBlogPosts() []Post {
	posts := []Post{}
	rows, err := db.Query("SELECT title, body, created FROM posts")
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(rows)
	for rows.Next() {
		p := Post{}
		err = rows.Scan(&p.Title, &p.Body, &p.Created)
		// fmt.Println(p)
		posts = append(posts, p)
	}
	return posts
}

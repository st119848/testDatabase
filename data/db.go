package data

import (
	_ "database/sql"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
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
func AddPost(p Post) error {
	stmt := "INSERT INTO posts (title, body, created) VALUES (?, ?, ?)"
	_, err := db.Exec(stmt, p.Title, p.Body, time.Now())
	return err
}
func (p Post) Save() error {
	stmt := "INSERT INTO posts (title, body, created) VALUES (?, ?, ?)"
	_, err := db.Exec(stmt, p.Title, p.Body, time.Now())
	return err
}
func PostById(id int64) Post {
	post := Post{}
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	sql, args, err := psql.
		Select("title, body, created").
		From("posts").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Get(&post, sql, args...)
	if err != nil {
		log.Fatalln(err)
	}

	return post
}

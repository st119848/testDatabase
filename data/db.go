package data

import (
	_ "database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init() {
	db = sqlx.MustConnect("mysql", "root:root@tcp(localhost)/post?parseTime=true")
}

package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbbook *sqlx.DB
var errbook error

type Book struct {
	ID     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func main() {
	initBookDB()
	query()
}

func query() {
	var books []Book
	err2 := dbbook.Select(&books, "select id,title,author,price from books where price > ?", 50)
	if err2 != nil {
		log.Fatal(err2)
	}
	for _, v := range books {
		fmt.Println(v)
	}
}

func initBookDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8"
	dbbook, errbook = sqlx.Connect("mysql", dsn)
	if errbook != nil {
		log.Fatal(errbook)
		return
	}

	dbbook.SetMaxOpenConns(30)
	dbbook.SetMaxIdleConns(10)
	return
}

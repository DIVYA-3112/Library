package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error
	dsn := "root:root@123@tcp(127.0.0.1:3306)/Library"
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	// defer DB.Close()

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("MySQL connection established")
}

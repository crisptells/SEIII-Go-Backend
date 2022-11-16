package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Db1() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:password@tcp(192.168.136.1:3306)/testdb")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	fmt.Println("Connected to database")
}

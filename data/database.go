package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
    idgames int `json:"id"`
    //Land1 string `json:"Land1"`
	//Land2 string `json:"Land2"`
}

func Db1() {
	fmt.Println("Go MySQL Tutorial")

	db, err := sql.Open("mysql", "root:LuisMaier@tcp(127.0.0.1:3306)/tippspiel_schema")
	if err != nil {
        fmt.Println(err.Error())
    }
	defer db.Close()
	fmt.Println("Connected to DB")

	pingErr := db.Ping()
    if pingErr != nil {
        fmt.Println(pingErr)
    }
    fmt.Println("Connected!")

	results, err := db.Query("SELECT idgames FROM games")
	fmt.Println("Got the results")
	

	if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

	for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.idgames)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        fmt.Println(tag.idgames)
    }

	defer db.Close()

	fmt.Println("Connected to database")
}

package sqldb

import(
    "database/sql"
    //"github.com/go-sql-driver/mysql"
	"fmt"
)

var DB *sql.DB

func Connect() {
	//Open connection to DB
	db, err := sql.Open("mysql", "root:LuisMaier@tcp(127.0.0.1:3306)/tippspiel_schema") //<name of the DB>
	if err != nil {
		panic(err.Error())
	}

	DB = db

	fmt.Println("Connected to DB")
}

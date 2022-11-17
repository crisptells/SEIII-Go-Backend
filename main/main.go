package main

import (
	"database/sql"
	data "example/user/Luis/data"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func getRoot(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "hello :)")

}

func main() {
	//Open connection to DB
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/tippspiel_db") //<name of the DB>
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	router := gin.Default()
	router.GET("/", getRoot)

	//Routes for users
	router.GET("/Users", data.GetAllUsers(db))
	router.POST("/AddUser", data.InsertUser(db))

	//Routes for games
	router.GET("/Games", data.GetAllGames(db))
	router.POST("/AddGame", data.InsertGame(db))
	router.Run("localhost:3333")
}

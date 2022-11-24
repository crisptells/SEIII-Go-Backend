package main

import (
	data "example/user/Luis/data"
	"example/user/Luis/globals"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func getRoot(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "hello :)")

}

func main() {
	sqldb.Connect()

	router := gin.Default()
	router.GET("/", getRoot)

	//Routes for users
	router.GET("/Users", data.GetAllUsers())
	router.POST("/AddUser", data.InsertUser())
	router.POST("/UpdateUser", data.UpdateUser())
	router.POST("/LoginUser", data.LoginUser())
	router.Run("localhost:3333")
}

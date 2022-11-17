package data

import (
	"database/sql"
	structure "example/user/Luis/structures"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Email   string
	Name    string
	Vorname string
	Geld    string
)

func InsertUser(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {
		//Create empty user
		var newUser structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&newUser); err != nil {
			panic(err)
		}
		insert, err := db.Prepare("INSERT INTO `users`(`email`,`passwort`,`name`,`vorname`, `geld`)VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		_, insertErr := insert.Exec(&newUser.Email, &newUser.Password, &newUser.Name, &newUser.FirstName, 0)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newUser)
	}
	return gin.HandlerFunc(fn)
}

func GetAllUsers(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		rows, err := db.Query("select * from users")
		if err != nil {
			fmt.Printf("Error: cant insert into users")
			panic(err.Error())
		}

		var users []structure.User
		for rows.Next() {
			var user structure.User
			if userErr := rows.Scan(&user.Email, &user.Password, &user.Name, &user.FirstName, &user.Cash); userErr != nil {
				log.Fatal(userErr)
			}
			users = append(users, user)
		}

		c.IndentedJSON(200, users)
	}

	return gin.HandlerFunc(fn)
}

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
		insert, err := db.Prepare("INSERT INTO `users`(`email`,`password`,`name`,`Vorname`, `Geld`)VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal(err)
		}

		_, insertErr := insert.Exec(&newUser.Email, &newUser.Password, &newUser.Name, &newUser.FirstName, &newUser.Cash)
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

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		//Create empty user
		var user structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&user); err != nil {
			panic(err)
		}
		_, err := db.Query("UPDATE users SET email = ?, password = ?, name = ?, vorname = ?, geld = ? WHERE email = ?", user.Email, user.Password, user.Name, user.FirstName, user.Cash, user.Email)
		if err != nil {
			fmt.Printf("Error: cant update user")
			panic(err.Error())
		}

		c.IndentedJSON(200, user)
	}

	return gin.HandlerFunc(fn)
}

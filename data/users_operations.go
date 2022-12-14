package data

import (
	//"database/sql"

	sqldb "example/user/Luis/globals"
	structure "example/user/Luis/structures"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Email    string
	Name     string
	Vorname  string
	Geld     string
	Subjects string
	Role     string
)

func InsertUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var newUser structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(400, "Failed to bind given values")
			panic(err)
		}
		//Check if email already exists
		rows, _ := sqldb.DB.Query("SELECT email FROM users WHERE email='" + newUser.Email + "'")

		if rows.Next() {
			c.IndentedJSON(400, "Email already exists")
			return
		}

		//Create user on the DB
		insert, err := sqldb.DB.Prepare("INSERT INTO `users`(`email`,`password`,`name`,`firstName`, `geld`)VALUES(?, ?, ?, ?, ?)")

		if err != nil {
			c.IndentedJSON(400, "cant create user in DB")
			panic(err)
		}

		_, insertErr := insert.Exec(&newUser.Email, &newUser.Password, &newUser.Name, &newUser.Vorname, &newUser.Geld)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(201, newUser)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := sqldb.DB.Query("select * from users")
		if err != nil {
			c.IndentedJSON(400, "cant find users")
			panic(err.Error())
		}

		var users []structure.User
		for rows.Next() {
			var user structure.User
			if userErr := rows.Scan(&user.Email, &user.Password, &user.Name, &user.Vorname, &user.Geld); userErr != nil {
				log.Fatal(userErr)
			}
			users = append(users, user)
		}

		c.IndentedJSON(200, users)
	}
}

func GetSpecificUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var user structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&user); err != nil {
			c.IndentedJSON(400, "wrong Email?")
			panic(err)
		}

		//get all emails and check if email is there, to prevent db error

		DBmails, err := sqldb.DB.Query("SELECT email FROM users")
		noMail := true
		for DBmails.Next() {
			var dbMail string
			if userErr := DBmails.Scan(&dbMail); userErr != nil {
				log.Fatal(userErr)
			}
			if user.Email == dbMail {
				noMail = false
			}
		}

		if noMail {
			c.IndentedJSON(400, "User not found")
			return
		}

		row, err := sqldb.DB.Query("SELECT * FROM users WHERE email = ?", user.Email)
		if err != nil {
			fmt.Printf("Error: cant find user")
			panic(err.Error())
		}
		if row == nil {
			c.IndentedJSON(404, "User not found")
			return
		}

		row.Next()
		if userErr := row.Scan(&user.Email, &user.Password, &user.Name, &user.Vorname, &user.Geld); userErr != nil {
			log.Fatal(userErr)
		}

		c.IndentedJSON(200, user)
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var user structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&user); err != nil {
			c.IndentedJSON(400, "wrong Email?")
			panic(err)
		}
		_, err := sqldb.DB.Query("UPDATE users SET email = ?, password = ?, name = ?, vorname = ?, geld = ? WHERE email = ?", user.Email, user.Password, user.Name, user.Vorname, user.Geld, user.Email)
		if err != nil {
			fmt.Printf("Error: cant update user")
			panic(err.Error())
		}

		c.IndentedJSON(200, user)
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var user structure.User
		var dataUser structure.User
		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&user); err != nil {
			panic(err)
		}
		row := sqldb.DB.QueryRow("SELECT * FROM users WHERE email = ?", user.Email)
		err := row.Scan(&dataUser.Email, &dataUser.Password)
		if err != nil {
			fmt.Printf("Error: cant get user")
			panic(err.Error())
		}

		if user.Password != dataUser.Password {
			c.IndentedJSON(406, "Wrong pw")
			return
		}

		c.IndentedJSON(200, user)
	}
}

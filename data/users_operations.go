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
	return func(c *gin.Context) {
		//Create empty user
		var newUser structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&newUser); err != nil {
			c.IndentedJSON(400, "Failed to bind given values")
			panic(err)
		}
		//Check if email already exists
		rows, err := db.Query("SELECT email FROM users")
		var mails []string
		for rows.Next() {
			var mail string
			if userErr := rows.Scan(&mail); userErr != nil {
				log.Fatal(userErr)
			}
			mails = append(mails, mail)
		}

		//Check if given mail is in all the mails from DB
		for _, mail := range mails {
			if mail == newUser.Email {
				c.IndentedJSON(400, "Email already exists")
				return
			}
		}

		//Create user on the DB
		insert, err := db.Prepare("INSERT INTO `users`(`email`,`password`,`name`,`Vorname`, `Geld`)VALUES(?, ?, ?, ?, ?)")

		if err != nil {
			panic(err)
			return
		}

		_, insertErr := insert.Exec(&newUser.Email, &newUser.Password, &newUser.Name, &newUser.FirstName, &newUser.Cash)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newUser)
	}
}

func GetAllUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("select * from users")
		if err != nil {
			c.IndentedJSON(400, "cant find games")
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
}

func UpdateUser(db *sql.DB) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		//Create empty user
		var user structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&user); err != nil {
			c.IndentedJSON(400, "wrong Email?")
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

func LoginUser(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var user structure.User
		var dataUser structure.User
		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&user); err != nil {
			panic(err)
		}
		row := db.QueryRow("SELECT * FROM users WHERE email = ?", user.Email)
		err := row.Scan(&dataUser.Email, &dataUser.Password, &dataUser.Name, &dataUser.FirstName, &dataUser.Cash)
		if err != nil {
			fmt.Printf("Error: cant get user")
			panic(err.Error())
		}

		if (user.Password != dataUser.Password) {
			c.IndentedJSON(406, "Wrong pw")
			return
		}

		c.IndentedJSON(200, user)
	}
}

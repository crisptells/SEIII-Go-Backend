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
		rows, _ := sqldb.DB.Query("SELECT email FROM users")
		var mails []string
		for rows.Next() {
			var mail string
			if userErr := rows.Scan(&mail); userErr != nil {
				fmt.Println(userErr.Error())
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
		insert, err := sqldb.DB.Prepare("INSERT INTO `users`(`email`,`password`,`name`,`firstName`, `role`, `subjects`)VALUES(?, ?, ?, ?, ?, ?)")

		if err != nil {
			c.IndentedJSON(400, "cant create user in DB")
			panic(err)
		}

		_, insertErr := insert.Exec(&newUser.Email, &newUser.Password, &newUser.Name, &newUser.FirstName, &newUser.Role, &newUser.Subjects)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newUser)
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := sqldb.DB.Query("select * from users")
		if err != nil {
			c.IndentedJSON(400, "cant find games")
			panic(err.Error())
		}

		var users []structure.User
		for rows.Next() {
			var user structure.User
			if userErr := rows.Scan(&user.Email, &user.Password, &user.Name, &user.FirstName, &user.Role, &user.Subjects); userErr != nil {
				log.Fatal(userErr)
			}
			users = append(users, user)
		}

		c.IndentedJSON(200, users)
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
		_, err := sqldb.DB.Query("UPDATE users SET email = ?, password = ?, name = ?, firstName = ?, role = ?, subjects = ? WHERE email = ?", user.Email, user.Password, user.Name, user.FirstName, user.Role, user.Subjects, user.Email)
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

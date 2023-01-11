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

func InsertExperience() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty Experience
		var newExperience structure.Experience

		//BindJSON to bind the received JSON to newExperience
		if err := c.BindJSON(&newExperience); err != nil {
			c.IndentedJSON(400, "Failed to bind given values")
			panic(err)
		}
		//Check if email already exists
		rows, err := sqldb.DB.Query("SELECT user_email FROM userexp WHERE user_email='" + newExperience.User_email + "'")

		if rows.Next() {
			c.IndentedJSON(400, "Exp for users email already exists")
			panic(err)
		}

		//Create Experience on the DB
		insert, err := sqldb.DB.Prepare("INSERT INTO `userexp`(`user_email`,`Math`,`German`,`English`, `Physics`, `Chemistry`, `Informatics`)VALUES(?, ?, ?, ?, ?, ?, ?)")

		if err != nil {
			c.IndentedJSON(400, "cant create user in DB")
			panic(err)
		}

		_, insertErr := insert.Exec(&newExperience.User_email, &newExperience.Math, &newExperience.German, &newExperience.English, &newExperience.Physics, &newExperience.Chemistry, &newExperience.Informatics)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newExperience)
	}
}

func GetAllExperiences() gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := sqldb.DB.Query("select * from userexp")
		if err != nil {
			c.IndentedJSON(400, "cant find experiences")
			panic(err.Error())
		}

		var experiences []structure.Experience
		for rows.Next() {
			var newExperience structure.Experience
			if userErr := rows.Scan(&newExperience.User_email, &newExperience.Math, &newExperience.German, &newExperience.English, &newExperience.Physics, &newExperience.Chemistry, &newExperience.Informatics); userErr != nil {
				log.Fatal(userErr)
			}
			experiences = append(experiences, newExperience)
		}

		c.IndentedJSON(200, experiences)
	}
}

func AddExperience() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty exp
		var experience structure.Experience
		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&experience); err != nil {
			c.IndentedJSON(400, "wrong")
			panic(err)
		}
		//Check if user has any prior xp
		rows, err := sqldb.DB.Query("select user_email from userexp")
		if err != nil {
			c.IndentedJSON(400, "cant find experiences")
			panic(err.Error())
		}
		userHasExp := false
		for rows.Next() {
			var dbMail string
			if userErr := rows.Scan(&dbMail); userErr != nil {
				log.Fatal(userErr)
			}
			if experience.User_email == dbMail {
				userHasExp = true
			}
		}

		if !userHasExp {
			_, err := sqldb.DB.Query("INSERT INTO `userexp`(`user_email`,`Math`,`German`,`English`, `Physics`, `Chemistry`, `Informatics`)VALUES(?, 0, 0, 0, 0, 0, 0)", experience.User_email)
			if err != nil {
				c.IndentedJSON(400, "cant initialize userexp")
				panic(err.Error())
			}
		} else {
			rows, err := sqldb.DB.Query("select * from userexp where user_email = ?", experience.User_email)
			if err != nil {
				c.IndentedJSON(401, "cant find users experiences")
				panic(err.Error())
			}
			var dbExperience structure.Experience
			rows.Next()
			if userErr := rows.Scan(&dbExperience.User_email, &dbExperience.Math, &dbExperience.German, &dbExperience.English, &dbExperience.Physics, &dbExperience.Chemistry, &dbExperience.Informatics); userErr != nil {
				log.Fatal(userErr)
			}
			experience.Math = dbExperience.Math + experience.Math
			experience.German = dbExperience.German + experience.German
			experience.English = dbExperience.English + experience.English
			experience.Physics = dbExperience.Physics + experience.Physics
			experience.Chemistry = dbExperience.Chemistry + experience.Chemistry
			experience.Informatics = dbExperience.Informatics + experience.Informatics

			_, err2 := sqldb.DB.Query("UPDATE userexp SET user_email = ?, Math = ?, German = ?, English = ?, Physics = ?, Chemistry = ?, Informatics = ? WHERE user_email = ?", experience.User_email, experience.Math, experience.German, experience.English, experience.Physics, experience.Chemistry, experience.Informatics, experience.User_email)
			if err2 != nil {
				fmt.Printf("Error: cant add experience")
				panic(err.Error())
			}
		}
		c.IndentedJSON(200, experience)
	}
}

func GetExperienceForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var user structure.User

		//BindJSON to bind the received JSON to newUser
		if err := c.BindJSON(&user); err != nil {
			c.IndentedJSON(400, "wrong Email?")
			panic(err)
		}
		experience, err := sqldb.DB.Query("SELECT * FROM userexp WHERE user_email='" + user.Email + "'")
		if err != nil {
			fmt.Printf("Error: cant get Exp for user")
			panic(err.Error())
		}
		var newExperience structure.Experience
		for experience.Next() {

			if userErr := experience.Scan(&newExperience.User_email, &newExperience.Math, &newExperience.German, &newExperience.English, &newExperience.Physics, &newExperience.Chemistry, &newExperience.Informatics); userErr != nil {
				log.Fatal(userErr)
			}
		}

		c.IndentedJSON(200, newExperience)
	}
}

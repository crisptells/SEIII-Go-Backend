package data

import (
	//"database/sql"
	sqldb "example/user/Luis/globals"
	structure "example/user/Luis/structures"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InsertTutoring() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty tutoring
		var newTutoring structure.Tutoring

		//BindJSON to bind the received JSON tonewTutoring
		if err := c.BindJSON(&newTutoring); err != nil {
			c.IndentedJSON(400, "Failed to bind given values")
			panic(err)
		}
		//Check if Tutoring already exists
		rows, _ := sqldb.DB.Query("SELECT tutoring_id FROM tutorings WHERE Tutor='" + newTutoring.Tutor + "' AND Subject='" + newTutoring.Subject + "'")

		if rows != nil {
			c.IndentedJSON(400, "Tutoring already exists")
			return
		}

		//Create Tutoring on the DB
		insert, err := sqldb.DB.Prepare("INSERT INTO `tutorings`(`Tutor`,`Subject`,`Students`,`MaxStudents`)VALUES(?, ?, ?, ?)")

		if err != nil {
			c.IndentedJSON(400, "error when inserting new Tutoring")
			panic(err)
		}

		_, insertErr := insert.Exec(&newTutoring.Tutoring_id, &newTutoring.Tutor, &newTutoring.Subject, &newTutoring.Students, &newTutoring.MaxStudents)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newTutoring)
	}
}

func GetAllTutorings() gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := sqldb.DB.Query("select * from tutorings")
		if err != nil {
			c.IndentedJSON(400, "cant find tutorings")
			panic(err.Error())
		}

		var tutorings []structure.Tutoring
		for rows.Next() {
			var tutoring structure.Tutoring
			if tutoringErr := rows.Scan(&tutoring.Tutoring_id, &tutoring.Tutor, &tutoring.Subject, &tutoring.Students, &tutoring.MaxStudents); tutoringErr != nil {
				log.Fatal(tutoringErr)
			}
			tutorings = append(tutorings, tutoring)
		}

		c.IndentedJSON(200, tutorings)
	}
}

func InsertUserTutoring() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var user_tutoring structure.User_tutoring
		//BindJSON to bind the received JSON to user
		if err := c.BindJSON(&user_tutoring); err != nil {
			c.IndentedJSON(400, "wrong Email?")
			panic(err)
		}
		//get all emails and check if email is there, to prevent db error
		DBmails, _ := sqldb.DB.Query("SELECT user_email FROM users_tutorings WHERE user_email = ?", user_tutoring.User_email)
		isEmailAlreadyThere := false
		for DBmails.Next() {
			var dbMail string
			if userErr := DBmails.Scan(&dbMail); userErr != nil {
				log.Fatal(userErr)
			}
			if user_tutoring.User_email == dbMail {
				isEmailAlreadyThere = true
			}
		}

		if isEmailAlreadyThere {
			//get all tutorings and check if tutoring_id is there, to prevent db error
			DBmails, _ := sqldb.DB.Query("SELECT tutoring_id FROM users_tutorings WHERE user_email = ?", user_tutoring.User_email)
			isTutoringAlreadyThere := false
			for DBmails.Next() {
				var dbMail int
				if userErr := DBmails.Scan(&dbMail); userErr != nil {
					log.Fatal(userErr)
				}
				intVar, _ := strconv.Atoi(user_tutoring.Tutoring_id)
				if intVar == dbMail {
					isTutoringAlreadyThere = true
				}
			}

			if isEmailAlreadyThere && isTutoringAlreadyThere {
				c.IndentedJSON(401, "Email and Tutoring is already there!")
				return
			}

		}
		//Insert new usertutoring
		//Create Tutoring on the DB

		insert, err := sqldb.DB.Prepare("INSERT INTO `users_tutorings`(`user_email`,`tutoring_id`)VALUES(?, ?)")

		if err != nil {
			c.IndentedJSON(400, "error when inserting new Tutoring")
			panic(err)
		}
		intTutoring_id, _ := strconv.Atoi(user_tutoring.Tutoring_id)
		_, insertErr := insert.Exec(&user_tutoring.User_email, &intTutoring_id)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, user_tutoring)
	}
}

func GetUserTutorings() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var user structure.User
		//BindJSON to bind the received JSON to user
		if err := c.BindJSON(&user); err != nil {
			c.IndentedJSON(400, "wrong Email?")
			panic(err)
		}
		fmt.Printf(user.Email)
		//get all emails and check if email is there, to prevent db error
		DBmails, _ := sqldb.DB.Query("SELECT user_email FROM users_tutorings")
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
			c.IndentedJSON(402, "User not found")
			return
		}

		row, err := sqldb.DB.Query("SELECT tutoring_id FROM users_tutorings WHERE user_email = ?", user.Email)
		if err != nil {
			fmt.Printf("Error: cant find user")
			panic(err.Error())
		}
		if row == nil {
			c.IndentedJSON(404, "No tutorings not found")
			return
		}

		//Get all tutoring_ids for the user
		var tutoring_ids []int
		for row.Next() {
			var tutoring_id int
			if userErr := row.Scan(&tutoring_id); userErr != nil {
				log.Fatal(userErr)
			}
			tutoring_ids = append(tutoring_ids, tutoring_id)
		}
		//Get all the Tutorings for the user
		var userstutorings []structure.Tutoring
		for index, _ := range tutoring_ids {
			tutorings_row, err := sqldb.DB.Query("SELECT * FROM tutorings WHERE tutoring_id = ?", tutoring_ids[index])
			if err != nil {
				fmt.Printf("Error: cant find user")
				panic(err.Error())
			}
			var tutoring structure.Tutoring
			tutorings_row.Next()
			if err := tutorings_row.Scan(&tutoring.Tutoring_id, &tutoring.Tutor, &tutoring.Subject, &tutoring.Students, &tutoring.MaxStudents); err != nil {
				log.Fatal(err)
			}
			userstutorings = append(userstutorings, tutoring)
		}

		c.IndentedJSON(200, userstutorings)
	}
}

func UpdateTutoring() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty tutoring
		var tutoring structure.Tutoring

		//BindJSON to bind the received JSON tonewTutoring
		if err := c.BindJSON(&tutoring); err != nil {
			c.IndentedJSON(400, "wrong Tutoring?")
			panic(err)
		}
		_, err := sqldb.DB.Query("UPDATE tutorings SET Tutor = ?, Subject = ?, Students = ?, MaxStudents = ? WHERE tutoring_id", tutoring.Tutor, tutoring.Subject, tutoring.Students, tutoring.MaxStudents, tutoring.Tutoring_id)
		if err != nil {
			c.IndentedJSON(400, "cant update tutoring")
			panic(err.Error())
		}

		c.IndentedJSON(200, tutoring)
	}
}

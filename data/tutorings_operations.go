package data

import (
	//"database/sql"
	structure "example/user/Luis/structures"
	"example/user/Luis/globals"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Tutor 	  		structure.User 	
	Subject   		string 
	Student   		structure.User 	
	MaxStudents 	string 	
	Appointments	[]structure.Appointment
)

func InsertTutoring() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var newTuroring structure.Tutoring

		//BindJSON to bind the received JSON to newTuroring
		if err := c.BindJSON(&newTuroring); err != nil {
			c.IndentedJSON(400, "Failed to bind given values")
			panic(err)
		}
		//Check if Tutoring already exists
		rows, err := sqldb.DB.Query("SELECT tutor, subject FROM tutorings")
		var keys []string
		for rows.Next() {
			var mail string
			if userErr := rows.Scan(&mail); userErr != nil {
				fmt.Printf(userErr.Error())
			}
			mails = append(mails, mail)
		}

		//Check if given mail is in all the mails from DB
		for _, mail := range mails {
			if mail == newTuroring.Email {
				c.IndentedJSON(400, "Email already exists")
				return
			}
		}

		//Create Tutoring on the DB
		insert, err := sqldb.DB.Prepare("INSERT INTO `tutorings`(`tutor`,`subject`,`students`,`maxStudents`, `appointments`)VALUES(?, ?, ?, ?, ?)")

		if err != nil {
			panic(err)
			return
		}

		_, insertErr := insert.Exec(&newTuroring.Tutor, &newTuroring.Subject, &newTuroring.Student, &newTuroring.MaxStudents, &newTuroring.Appointments)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newTuroring)
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
			if tutoringErr := rows.Scan(&tutoring.Tutor, &tutoring.Subject, &tutoring.Student, &tutoring.MaxStudents, &tutoring.Appointments); userErr != nil {
				log.Fatal(tutoringErr)
			}
			users = append(users, user)
		}

		c.IndentedJSON(200, users)
	}
}

func UpdateTutoring() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty tutoring
		var tutoring structure.Tutoring

		//BindJSON to bind the received JSON to newTuroring
		if err := c.BindJSON(&tutoring); err != nil {
			c.IndentedJSON(400, "wrong Email?")
			panic(err)
			return
		}
		_, err := sqldb.DB.Query("UPDATE users SET email = ?, password = ?, name = ?, firstName = ?, role = ?, subjects = ? WHERE email = ?", user.Email, user.Password, user.Name, user.FirstName, user.Role, user.Subjects, user.Email)
		if err != nil {
			fmt.Printf("Error: cant update user")
			panic(err.Error())
		}

		c.IndentedJSON(200, user)
	}
}
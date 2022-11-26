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
	Tutor        structure.User
	Subject      string
	Student      structure.User
	MaxStudents  int
	Appointments []structure.Appointment
)

type tutoringKey struct {
	Tutor   string
	Subject string
}

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

		if err == nil {
			var keys []tutoringKey
			for rows.Next() {
				var key tutoringKey
				if userErr := rows.Scan(&key.Tutor, &key.Subject); userErr != nil {
					fmt.Println(userErr.Error())
				}
				keys = append(keys, key)
			}

			//Check if given Tutor and Subject combo is already there
			for _, checkKey := range keys {
				if checkKey.Tutor == newTuroring.Tutor && checkKey.Subject == newTuroring.Subject {
					c.IndentedJSON(400, "Email already exists")
					return
				}
			}
		}

		//Create Tutoring on the DB
		insert, err := sqldb.DB.Prepare("INSERT INTO `tutorings`(`Tutor`,`Subject`,`Students`,`MaxStudents`, `Appointments`)VALUES(?, ?, ?, ?, ?)")

		if err != nil {
			panic(err)
			return
		}

		_, insertErr := insert.Exec(&newTuroring.Tutor, &newTuroring.Subject, &newTuroring.Students, &newTuroring.MaxStudents, &newTuroring.Appointments)
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
			if tutoringErr := rows.Scan(&tutoring.Tutor, &tutoring.Subject, &tutoring.Students, &tutoring.MaxStudents, &tutoring.Appointments); tutoringErr != nil {
				log.Fatal(tutoringErr)
			}
			tutorings = append(tutorings, tutoring)
		}

		c.IndentedJSON(200, tutorings)
	}
}

func UpdateTutoring() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty tutoring
		var tutoring structure.Tutoring

		//BindJSON to bind the received JSON to newTuroring
		if err := c.BindJSON(&tutoring); err != nil {
			c.IndentedJSON(400, "wrong Tutoring?")
			panic(err)
			return
		}
		_, err := sqldb.DB.Query("UPDATE tutorings SET Tutor = ?, Subject = ?, Students = ?, MaxStudents = ?, Appointments = ? WHERE Tutor = ? AND Subject = ?", tutoring.Tutor, tutoring.Subject, tutoring.Students, tutoring.MaxStudents, tutoring.Appointments, tutoring.Tutor, tutoring.Subject)
		if err != nil {
			fmt.Printf("Error: cant update user")
			panic(err.Error())
		}

		c.IndentedJSON(200, tutoring)
	}
}

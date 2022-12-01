package data

import (
	//"database/sql"
	sqldb "example/user/Luis/globals"
	structure "example/user/Luis/structures"
	"log"

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

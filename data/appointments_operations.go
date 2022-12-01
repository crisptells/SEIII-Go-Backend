package data

import (
	//"database/sql"
	sqldb "example/user/Luis/globals"
	structure "example/user/Luis/structures"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func InserAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty user
		var newAppointment structure.Appointment

		//BindJSON to bind the received JSON to newAppointment
		if err := c.BindJSON(&newAppointment); err != nil {
			c.IndentedJSON(400, "Failed to bind given values")
			panic(err)
		}
		//Check if Appointment already exists
		rows, _ := sqldb.DB.Query("SELECT appointment_id FROM appointments WHERE Date='" + newAppointment.Date + "' AND tutoring_id='" + newAppointment.Tutoring_id + "'")

		if rows.Next() {
			c.IndentedJSON(400, "Appointment for date and tutoring already exists")
			return
		}

		//Create Tutoring on the DB
		insert, err := sqldb.DB.Prepare("INSERT INTO `appointments`(`Date`,`Duration`,`tutoring_id`)VALUES(?, ?, ?)")

		if err != nil {
			c.IndentedJSON(400, "error when inserting new Appointment")
			panic(err)
		}

		_, insertErr := insert.Exec(&newAppointment.Date, &newAppointment.Duration, &newAppointment.Tutoring_id)
		if insertErr != nil {
			log.Fatal(insertErr)
		}

		c.IndentedJSON(200, newAppointment)
	}
}

func GetAllAppointments() gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := sqldb.DB.Query("select * from appointments")
		if err != nil {
			c.IndentedJSON(400, "cant find appointments")
			panic(err.Error())
		}

		if rows == nil {
			c.IndentedJSON(400, "no appointments in DB")
			panic(err.Error())
		}

		var appointments []structure.Appointment
		for rows.Next() {
			var appointment structure.Appointment
			if tutoringErr := rows.Scan(&appointment.Appointment_id, &appointment.Date, &appointment.Duration, &appointment.Tutoring_id); tutoringErr != nil {
				log.Fatal(tutoringErr)
			}
			appointments = append(appointments, appointment)
		}

		c.IndentedJSON(200, appointments)
	}
}

func GetTutoringAppointments() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty appointment
		var newAppointment structure.Appointment

		//BindJSON to bind the received JSON to newAppointment
		if err := c.BindJSON(&newAppointment); err != nil {
			c.IndentedJSON(400, "Failed to bind given values")
			panic(err)
		}

		rows, err := sqldb.DB.Query("SELECT * FROM appointments WHERE tutoring_id='" + newAppointment.Tutoring_id + "'")

		if err != nil {
			c.IndentedJSON(400, "Appointment for date and tutoring already exists")
			return
		}

		var appointments []structure.Appointment
		for rows.Next() {
			var appointment structure.Appointment
			if tutoringErr := rows.Scan(&appointment.Appointment_id, &appointment.Date, &appointment.Duration, &appointment.Tutoring_id); tutoringErr != nil {
				log.Fatal(tutoringErr)
			}
			appointments = append(appointments, appointment)
		}

		c.IndentedJSON(200, appointments)

	}
}

func UpdateAppointment() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Create empty appointment
		var appointment structure.Appointment

		//BindJSON to bind the received JSON to newAppointment
		if err := c.BindJSON(&appointment); err != nil {
			c.IndentedJSON(400, "wrong Tutoring?")
			panic(err)
		}

		//Check if appointment_id is in DB
		rows, _ := sqldb.DB.Query("SELECT * FROM appointments WHERE tutoring_id='" + appointment.Tutoring_id + "'")

		if !rows.Next() {
			c.IndentedJSON(400, "Appointment does not exist")
			return
		}

		_, err := sqldb.DB.Query("UPDATE appointments SET Date = ?, Duration = ?, tutoring_id = ? WHERE appointment_id = ?", appointment.Date, appointment.Duration, appointment.Tutoring_id, appointment.Appointment_id)
		if err != nil {
			c.IndentedJSON(400, "cant update appointment")
			panic(err.Error())
		}

		c.IndentedJSON(200, appointment)
	}
}

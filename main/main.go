package main

import (
	data "example/user/Luis/data"
	sqldb "example/user/Luis/globals"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func getRoot(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "hello :)")

}

func main() {
	sqldb.Connect()

	router := gin.Default()
	router.GET("/", getRoot)

	//Routes for users
	router.GET("/Users", data.GetAllUsers())
	router.POST("/AddUser", data.InsertUser())
	router.POST("/UpdateUser", data.UpdateUser())
	router.POST("/LoginUser", data.LoginUser())

	//Routes for Tutorings
	router.GET("/Tutorings", data.GetAllTutorings())
	router.POST("/AddTutoring", data.InsertTutoring())
	router.POST("/UpdateTutoring", data.UpdateTutoring())

	//Routes for Appointments
	router.GET("/Appointments", data.GetAllAppointments())
	router.POST("/AddAppointment", data.InserAppointment())
	router.POST("/UpdateAppointment", data.UpdateAppointment())
	router.GET("/TutoringAppointments", data.GetTutoringAppointments())

	//Start service
	router.Run("localhost:3333")
}

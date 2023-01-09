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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	sqldb.Connect()
	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/", getRoot)

	//Routes for users
	router.GET("/Users", data.GetAllUsers())
	router.POST("/User", data.GetSpecificUser())
	router.POST("/AddUser", data.InsertUser())
	router.POST("/UpdateUser", data.UpdateUser())
	router.POST("/LoginUser", data.LoginUser())

	//Routes for Tutorings
	router.GET("/Tutorings", data.GetAllTutorings())
	router.POST("/AddTutoring", data.InsertTutoring())
	router.POST("/UpdateTutoring", data.UpdateTutoring())

	//Routes for users_tutorings
	router.POST("/AddUserTutoring", data.InsertUserTutoring())
	router.POST("/GetUsersTutorings", data.GetUserTutorings())

	//Routes for Appointments
	router.GET("/Appointments", data.GetAllAppointments())
	router.POST("/AddAppointment", data.InserAppointment())
	router.POST("/UpdateAppointment", data.UpdateAppointment())
	router.GET("/TutoringAppointments", data.GetTutoringAppointments())

	//Routes for Experiences
	router.GET("/Experiences", data.GetAllExperiences())
	router.POST("/AddExp", data.InsertExperience())
	router.POST("/CountUpExp", data.AddExperience())
	router.POST("/UserExp", data.GetExperienceForUser())

	//Start service
	router.Run("localhost:3333")
}

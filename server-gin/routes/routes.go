package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samuelreboucas07/api-go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowStudents)
	r.GET("/student/cpf/:cpf", controllers.SearchStudentByCpf)
	r.GET("/student/:id", controllers.ShowStudentById)
	r.DELETE("/student/:id", controllers.DeleteStudentById)
	r.PATCH("/student/:id", controllers.EditStudentById)
	r.GET("/:name", controllers.Hello)
	r.POST("/students", controllers.CreateStudent)
	r.Run()
}

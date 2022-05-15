package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samuelreboucas07/api-go-gin/database"
	"github.com/samuelreboucas07/api-go-gin/models"
)

func ShowStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func ShowStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Aluno não encontrado",
		})
		return
	}

	c.JSON(200, student)
}

func DeleteStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.Delete(&student, id)

	c.JSON(200, gin.H{
		"Message": "Estudante removido com sucesso",
	})
}

func EditStudentById(c *gin.Context) {
	id := c.Params.ByName("id")
	var student models.Student
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func SearchStudentByCpf(c *gin.Context) {
	cpf := c.Param("cpf")
	var student models.Student
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Message": "Estudante não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func Hello(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"Message": "Olá " + name + ", tudo bem?",
	})
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateStudent(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

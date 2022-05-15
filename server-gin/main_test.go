package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/samuelreboucas07/api-go-gin/controllers"
	"github.com/samuelreboucas07/api-go-gin/database"
	"github.com/samuelreboucas07/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func setupTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Teste", CPF: "12345678912", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestVerifyStatusCodeHello(t *testing.T) { //Função deve iniciar com a palavra Test
	r := setupTest()
	r.GET("/:name", controllers.Hello)
	request, _ := http.NewRequest("GET", "/samuel", nil)
	response := httptest.NewRecorder() //Simula um response

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Deveriam ser iguais")

	mockResponse := `{"Message":"Olá samuel, tudo bem?"}`
	responseBody, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(responseBody))
}

func TestGellAll(t *testing.T) {
	database.ConnectDatabase()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := setupTest()
	r.GET("/students", controllers.ShowStudents)
	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestSearchById(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := setupTest()
	r.GET("/student/cpf/:cpf", controllers.SearchStudentByCpf)
	request, _ := http.NewRequest("GET", "/student/cpf/12345678912", nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetById(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := setupTest()
	r.GET("/student/:id", controllers.ShowStudentById)
	path := "/student/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Teste", studentMock.Name)
	assert.Equal(t, "12345678912", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
}

func TestDeleteById(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()

	r := setupTest()
	r.DELETE("/student/:id", controllers.DeleteStudentById)
	path := "/student/" + strconv.Itoa(ID)
	request, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestEditById(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := setupTest()
	r.PATCH("/student/:id", controllers.EditStudentById)
	path := "/student/" + strconv.Itoa(ID)
	newStudent := models.Student{Name: "Teste novo", CPF: "12345678912", RG: "123456789"}
	newStudentJson, _ := json.Marshal(newStudent)
	request, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(newStudentJson))
	response := httptest.NewRecorder()

	r.ServeHTTP(response, request)

	var studentMock models.Student
	json.Unmarshal(response.Body.Bytes(), &studentMock)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, "Teste novo", studentMock.Name)
	assert.Equal(t, "12345678912", studentMock.CPF)
	assert.Equal(t, "123456789", studentMock.RG)
}

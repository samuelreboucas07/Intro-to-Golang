package main

import (
	"api-rest/database"
	"api-rest/models"
	"api-rest/routes"
	"fmt"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "samuel", Story: "Teste"},
		{Id: 2, Name: "carlos", Story: "Teste carlos"},
	}
	database.ConnectDatabase()
	fmt.Println("Iniciando servidor Rest em Go")
	routes.HandleRequest()
}

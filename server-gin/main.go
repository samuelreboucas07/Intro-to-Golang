package main

import (
	"github.com/samuelreboucas07/api-go-gin/database"
	"github.com/samuelreboucas07/api-go-gin/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequests()
}

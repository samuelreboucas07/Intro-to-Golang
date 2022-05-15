package database

import (
	"log"

	"github.com/samuelreboucas07/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connect := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(connect))

	if err != nil {
		log.Panic("Erro ao conectar como banco de dados")
	}
	DB.AutoMigrate(&models.Student{})
}

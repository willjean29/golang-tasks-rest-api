package db

import (
	taskModels "app/src/modules/tasks/infra/data/gorm/models"
	userModels "app/src/modules/users/infra/data/gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS string = "host=localhost user=admin password=admin dbname=tasks port=5432"
var DB *gorm.DB

func GormConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected successfully")
	}
}

func GormSyncDatabase() {
	GormConnection()
	DB.AutoMigrate(&taskModels.Task{}, &userModels.User{})
}

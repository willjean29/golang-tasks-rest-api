package db

import (
	"app/src/config/plugins"
	taskModels "app/src/modules/tasks/infra/data/gorm/models"
	userModels "app/src/modules/users/infra/data/gorm/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
	plugins.Envs.DBHost, plugins.Envs.DBUser, plugins.Envs.DBPassword, plugins.Envs.DBName, plugins.Envs.DBPort)
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

package database

import (
	"producer/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsnFormat := "host=localhost user=macbook password=root dbname=db_authentication port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsnFormat), &gorm.Config{})

	if err != nil {
		panic("Error fail connected database !")
	}

	DB = database

	database.AutoMigrate(&model.Policy{})
}

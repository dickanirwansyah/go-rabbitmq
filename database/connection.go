package database

import (
	"producer/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsnFormat := "host=localhost user=golang password=golang dbname=db_go port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsnFormat), &gorm.Config{})

	if err != nil {
		panic("Error fail connected database !")
	}

	DB = database

	database.AutoMigrate(
		&model.Policy{},
		&model.Roles{},
		&model.Account{},
		&model.PublishClaim{},
		&model.Permissions{},
		&model.PermissionsRoles{},
	)
}

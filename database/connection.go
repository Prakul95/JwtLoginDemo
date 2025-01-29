package database

import (
	"loginApp/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=rijul password=Pearcesquare dbname= golang_testdb port=5434 sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error)
	}
	DB = connection
	connection.AutoMigrate(&models.User{})

}

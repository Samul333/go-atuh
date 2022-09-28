package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDatabaseInstance() error {
	dsn := "host=localhost user=postgres password=admin dbname=go-auth port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Successfully connected to database")
	DB = db

	return nil
}

func SetupDatabaseModels() error {

	dbModels := []interface{}{&User{}}

	err := DB.AutoMigrate(dbModels...)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

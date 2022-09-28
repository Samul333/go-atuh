package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/samul333/go-auth/database"
)

func main() {
	app := fiber.New()

	SetupRoutes(app)

	err := database.NewDatabaseInstance()

	if err != nil {
		fmt.Println("Error connected to database")
		panic(err)
	}

	err = database.SetupDatabaseModels()

	if err != nil {
		fmt.Println("Error migrating models")
		panic(err)
	}

	fmt.Println(database.DB)

	app.Listen(":3000")

}

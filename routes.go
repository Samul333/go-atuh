package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samul333/go-auth/controllers"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	controllers.SetupUserRoutes(api)

}

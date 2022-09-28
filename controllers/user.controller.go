package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samul333/go-auth/handler"
	"github.com/samul333/go-auth/middleware"
)

func SetupUserRoutes(api fiber.Router) {

	userRoute := api.Group("/users")

	userRoute.Get("/", middleware.AuthRequired(), handler.UserHandlerTest)
	userRoute.Post("/register", handler.UserRegistrationHandler)
	userRoute.Post("/login", handler.UserLogin)

}

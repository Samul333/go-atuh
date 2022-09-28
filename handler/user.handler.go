package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/samul333/go-auth/database"
)

func UserHandlerTest(c *fiber.Ctx) error {
	users := []database.User{}
	database.DB.Find(&users)
	return c.JSON(users)

}

func UserRegistrationHandler(c *fiber.Ctx) error {
	jwtSecret := "supersecretpasswrod"

	user := database.User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	hashedPassword, err := HashPasswords(user.Password)

	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
		})
	}

	user.Password = hashedPassword

	err = database.DB.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
		})
	}

	s, err := GenerateJwtToken(&user, jwtSecret)

	if err != nil {
		return c.SendString("Error")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
	})

}

func UserLogin(c *fiber.Ctx) error {

	user := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
		})
	}

	curUser := database.User{}

	result := database.DB.Where(&database.User{Username: user.Username}).First(&curUser)

	if result.Error != nil {
		fmt.Println(result.Error)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
		})
	}

	validUser := DecodePassword(curUser.Password, user.Password)

	if !validUser {
		fmt.Println("Password is not valid")
		return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"success": false,
		})
	}

	s, err := GenerateJwtToken(&curUser, "supersecretpassword")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"token":   s,
	})

}

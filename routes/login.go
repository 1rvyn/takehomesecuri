package routes

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	success := true
	return c.JSON(fiber.Map{
		"message": "Login successful",
		"success": success,
		"user":    "user",
	})
}

package main

import (
	"fmt"

	"github.com/1rvyn/takehomesecuri/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	setupRoutes(app)

	fmt.Println("Server is running on port 3000")

	app.Listen(":3000")

}

func setupRoutes(app *fiber.App) {
	app.Post("/login", routes.Login)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}

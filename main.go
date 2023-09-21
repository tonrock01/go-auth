package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Simple GET handler
	app.Get("/api/list", func(c *fiber.Ctx) error {
		return c.SendString("I'm a GET request!")
	})

	// Simple POST handler
	app.Post("/api/register", func(c *fiber.Ctx) error {
		return c.SendString("I'm a POST request!")
	})

	app.Listen(":8080")
}

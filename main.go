package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/tonrock01/go-test-auth/db"
	"github.com/tonrock01/go-test-auth/router"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loadinf .env file")
	}

	app := fiber.New()
	app.Use(cors.New())

	db.InitDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/getcurrent", router.GetCurrentUser)
	app.Post("/signup", router.SignUpUser)
	app.Post("/signin", router.SignInUser)

	app.Listen(":8080")
}

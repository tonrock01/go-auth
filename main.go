package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/tonrock01/go-test-auth/db"
	"github.com/tonrock01/go-test-auth/handler"
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

	app.Get("/getcurrent", handler.GetCurrentUser)
	app.Post("/signup", handler.SignUpUser)
	app.Post("/signin", handler.SignInUser)

	log.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
}

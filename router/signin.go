package router

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tonrock01/go-test-auth/db"
	"github.com/tonrock01/go-test-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func SignInUser(c *fiber.Ctx) error {
	req := models.SignInRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	//Check field
	if req.Username == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Some field is empty"})
	}

	//Check user in database
	checkUser := models.User{}
	db.Db.Where("username = ?", req.Username).First(&checkUser)
	if checkUser.ID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "User dose not exist"})
	}

	//Check password
	err := bcrypt.CompareHashAndPassword([]byte(checkUser.Password), []byte(req.Password))
	if err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": checkUser.Username,
			"exp":      time.Now().Add(time.Minute * 1).Unix(),
		})
		t, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Login Success", "token": t})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Incorrect password"})
	}
}

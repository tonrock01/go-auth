package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tonrock01/go-test-auth/db"
	"github.com/tonrock01/go-test-auth/models"
)

func GetCurrentUser(c *fiber.Ctx) error {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	user := models.User{}
	db.Db.Where("username = ?", claims["username"].(string)).First(&user)

	responseUser := models.ResponseUser{
		Username:  user.Username,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "user": responseUser})
}

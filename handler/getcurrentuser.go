package handler

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tonrock01/go-test-auth/db"
	"github.com/tonrock01/go-test-auth/models"
)

func GetCurrentUser(c *fiber.Ctx) error {
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "Failed", "message": "Can not get current user"})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("invalidate token: %v", err)})
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := models.User{}
		db.Db.Where("username = ?", claims["username"]).First(&user)

		responseUser := models.ResponseUser{
			Username:  user.Username,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "Success", "user": responseUser})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "Failed", "message": "No info"})
	}
}

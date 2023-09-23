package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tonrock01/go-test-auth/db"
	"github.com/tonrock01/go-test-auth/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
	req := models.SignUpRequest{}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}

	//Check field
	if req.Username == "" || req.Password == "" || req.Firstname == "" || req.Lastname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid SignUp"})
	}

	//Check user in database
	checkUser := models.User{}
	db.Db.Where("username = ?", req.Username).First(&checkUser)
	if checkUser.ID > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "This username already used"})
	}

	//Hash password
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Username:  req.Username,
		Password:  string(hashpassword),
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
	}

	db.Db.Create(&user)
	if user.ID > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "SignUp Success"})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "SignUp Failed"})
	}
}

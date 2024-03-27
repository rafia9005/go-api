package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafia9005/go-api/database"
	"github.com/rafia9005/go-api/middleware"
	"github.com/rafia9005/go-api/model/entity"
	"github.com/rafia9005/go-api/model/request"
	"github.com/rafia9005/go-api/utils"
)

func Login(c *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := c.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
	}

	var users entity.Users

	err := database.DB.Debug().First(&users, "email = ?", loginRequest.Email).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = users.Name
	claims["email"] = users.Email
	claims["exp"] = time.Now().Add(time.Minute * 2).Unix()
	claims["role"] = "user"

	if users.Role == "admin" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Worng credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func Register(c *fiber.Ctx) error {
	users := new(request.RegisterRequest)
	if err := c.BodyParser(users); err != nil {
		return err
	}
	validate := validator.New()
	errValidate := validate.Struct(users)

	if errValidate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed required",
			"error":   errValidate.Error(),
		})
	}

	hashedPassword, err := middleware.HashPassword(users.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	newUsers := entity.Users{
		Name:     users.Name,
		Email:    users.Email,
		Password: hashedPassword,
		Role:     "user",
	}

	errCreate := database.DB.Create(&newUsers).Error
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store users",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success register",
	})
}

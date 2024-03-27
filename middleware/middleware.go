package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rafia9005/go-api/utils"
	"golang.org/x/crypto/bcrypt"
)

func Auth(c *fiber.Ctx) error {
	token := c.Get("x-token")

	if token == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	_, err := utils.VerifyToken(token)

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	c.Locals("usersInfo", claims)
	c.Locals("role", claims["role"])

	return c.Next()
}

func AdminRole(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role == "user" {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	return c.Next()
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

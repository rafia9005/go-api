package handler

import "github.com/gofiber/fiber/v2"

func Welcome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{
		"status":  "succes",
		"version": "1.0.0",
		"message": "API is working right now!",
	})
}

package handler

import "github.com/gofiber/fiber/v2"

func IndexProduct(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "testing",
	})
}

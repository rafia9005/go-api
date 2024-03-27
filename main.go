package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafia9005/go-api/database"
	"github.com/rafia9005/go-api/routes"
)

func main() {
	app := fiber.New()

	database.Connect()

	routes.AutoMigrate()

	routes.SetupRouter(app)

	app.Listen(":8000")
}

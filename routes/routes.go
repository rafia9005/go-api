package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafia9005/go-api/handler"
	auth "github.com/rafia9005/go-api/handler/Auth"
	"github.com/rafia9005/go-api/middleware"
	"github.com/rafia9005/go-api/model/entity"
)

var Admin = middleware.AdminRole
var Auth = middleware.Auth

func SetupRouter(app *fiber.App) {
	app.Get("/users", Auth, Admin, handler.IndexUsers)
	app.Post("/users", Auth, handler.CreateUsers)
	app.Get("/users/:id", Auth, handler.ShowUsers)
	app.Delete("/users/:id", Auth, handler.DeleteUsers)
	app.Put("/users/:id", Auth, handler.UpdateUser)
	app.Get("/product", Auth, handler.IndexProduct)

	// Endpoint untuk login
	app.Post("/login", auth.Login)
	app.Post("/register", auth.Register)

	// Endpoint untuk static file (jika diperlukan)
	app.Static("/", "./public")
}

func AutoMigrate() {
	RunMigrate(&entity.Users{})
	RunMigrate(&entity.Product{})
}

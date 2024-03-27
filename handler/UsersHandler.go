package handler

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rafia9005/go-api/database"
	"github.com/rafia9005/go-api/middleware"
	"github.com/rafia9005/go-api/model/entity"
	"github.com/rafia9005/go-api/model/request"
)

func IndexUsers(c *fiber.Ctx) error {
	var users []entity.Users
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	var userResponses []request.UserShowResponse

	for _, users := range users {
		userResponse := request.UserShowResponse{
			ID:    users.ID,
			Name:  users.Name,
			Email: users.Email,
		}
		userResponses = append(userResponses, userResponse)
	}

	return c.JSON(fiber.Map{
		"data": userResponses,
	})
}

func ShowUsers(c *fiber.Ctx) error {
	usersId := c.Params("id")

	var users entity.Users

	result := database.DB.First(&users, usersId)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	userResponse := request.UserShowResponse{
		ID:    users.ID,
		Name:  users.Name,
		Email: users.Email,
	}

	return c.JSON(fiber.Map{
		"data": userResponse,
	})
}

func DeleteUsers(c *fiber.Ctx) error {
	userID := c.Params("id")

	result := database.DB.Delete(&entity.Users{}, userID)
	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})
}

func CreateUsers(c *fiber.Ctx) error {
	users := new(request.UsersCreateRequest)
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
	}

	errCreate := database.DB.Create(&newUsers).Error
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "failed to store users",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"data":    newUsers,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	usersRequest := new(request.UsersUpdateRequest)
	if err := c.BodyParser(usersRequest); err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	usersId := c.Params("id")
	var users entity.Users
	err := database.DB.First(&users, "id = ?", usersId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "users not found",
		})
	}

	if usersRequest.Name != "" {
		users.Name = usersRequest.Name
	}
	if usersRequest.Email != "" {
		users.Email = usersRequest.Email
	}

	errUpdate := database.DB.Save(&users).Error
	if errUpdate != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	responseData := request.UserShowResponse{
		Name:  users.Name,
		Email: users.Email,
	}

	return c.JSON(fiber.Map{
		"message": "succes",
		"data":    responseData,
	})
}

func UserTesting() {

}

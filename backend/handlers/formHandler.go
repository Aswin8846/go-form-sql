package handlers

import (
	"project/database"

	"github.com/gofiber/fiber/v2"
)

func SubmitFormHandler(c *fiber.Ctx) error {
	// Parse request body
	var data struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		Comments string `json:"comments"`
	}

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid input",
		})
	}

	// Insert into database
	_, err := database.DB.Exec(
		"INSERT INTO users (name, email, phone, address, comments) VALUES (?, ?, ?, ?, ?)",
		data.Name, data.Email, data.Phone, data.Address, data.Comments,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to insert data",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Data submitted successfully",
	})
}


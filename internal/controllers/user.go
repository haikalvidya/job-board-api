package controllers

import (
	"job-board-api/cmd"
	"job-board-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func UserMe(c *fiber.Ctx) error {
	store := cmd.Http.Session.Get(c)
	userID := store.Get("user_id")
	if userID == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User Not Logged In",
			"data":    "User Not Logged In",
		})
	}

	user, err := models.GetUserById(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User Not Logged In",
			"data":    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "User Found",
		"data":    user,
	})
}

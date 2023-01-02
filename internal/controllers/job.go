package controllers

import (
	"job-board-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func JobRequest(c *fiber.Ctx) error {
	// Get Query Params
	paramsJob := models.JobParams{
		Description: c.Query("description"),
		Location:    c.Query("location"),
		FullTime:    c.Query("full_time"),
	}
	// get path params job id
	paramsJob.JobID = c.Params("jobId")
	jobs, err := models.GetJobs(paramsJob)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Jobs Not Found",
			"data":    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Success Get Jobs",
		"data":    jobs,
	})
}

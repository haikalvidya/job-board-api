package routers

import (
	"job-board-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func JobRouter(App fiber.Router) {
	App.Get("", controllers.JobRequest)
	App.Get("/:jobId", controllers.JobRequest)
}

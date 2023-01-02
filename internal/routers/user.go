package routers

import (
	"job-board-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(App fiber.Router) {
	App.Get("/me", controllers.UserMe)
}

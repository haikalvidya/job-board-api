package routers

import (
	"job-board-api/internal/routers/middleware"

	"github.com/gofiber/fiber/v2"
)

func LoadRoutes(App *fiber.App) {
	AuthRouter(App.Group("/api"))
	UserRouter(App.Group("/api/user").Use(middleware.Auth()))
	JobRouter(App.Group("/api/job").Use(middleware.Auth()))
}

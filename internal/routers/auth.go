package routers

import (
	"job-board-api/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(App fiber.Router) {
	App.Post("/register", controllers.ValidateRegisterUser, controllers.RegisterUser)
	App.Post("/login", controllers.ValidateLoginUser, controllers.LoginPost)
}

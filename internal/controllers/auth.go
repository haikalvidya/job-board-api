package controllers

import (
	"job-board-api/cmd"
	"job-board-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func RegisterUser(c *fiber.Ctx) error {
	register := c.Locals("register").(models.RegisterUser)
	user, err := register.Signup()
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "User Already Exists",
			"data":    err.Error(),
		})
	}
	session := cmd.Http.Session.Get(c)
	session.Set("user_id", user.ID)
	session.Save()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "User Created Successfully",
	})
}

// midleware register user
func ValidateRegisterUser(c *fiber.Ctx) error {
	var register models.RegisterUser
	if err := c.BodyParser(&register); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid JSON",
			"data":    err.Error(),
		})
	}

	v := validate.Struct(register)
	if !v.Validate() {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Validation Error",
			"data":    v.Errors.One(),
		})
	}
	c.Locals("register", register)
	return c.Next()
}

func LoginPost(c *fiber.Ctx) error { //nolint:wsl
	user := c.Locals("user").(*models.User)
	token, err := models.Login(c, user.ID, cmd.Http.Jwt.Secret) //nolint:wsl
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   true,
			"message": "Login Error",
			"data":    err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "Login Success",
		"data":    token,
	})
}

func ValidateLoginUser(c *fiber.Ctx) error {
	var login models.LoginUser
	if err := c.BodyParser(&login); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid JSON",
			"data":    err.Error(),
		})
	}
	v := validate.Struct(login)
	if !v.Validate() {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid JSON",
		})
	}

	user, err := models.GetUserByUsername(login.Username)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "invalid Username or Password",
			"data":    err.Error(),
		})
	}
	match, err := cmd.Http.Hash.Match(login.Password, user.Password)
	if !match {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "invalid Username or Password",
			"data":    v.Errors.One(),
		})
	}

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error":   true,
			"message": "Validation Error",
			"data":    v.Errors.One(),
		})
	}
	c.Locals("user", user)
	return c.Next()
}

func Logout(c *fiber.Ctx) error {
	store := cmd.Http.Session.Get(c)
	store.Delete("user_id")
	store.Delete("user_token")
	err := store.Save()
	if err != nil {
		panic(err)
	}
	c.ClearCookie()
	return nil
}

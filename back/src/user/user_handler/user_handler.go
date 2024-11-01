package user_handler

import (
	"back/src/user/user_models"
	"back/src/user/user_services"

	"regexp"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *user_services.UserService
}

func NewUserHandler(userService *user_services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name parameter is required"})
	}

	user, err := h.userService.GetGraphUser(name)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(200).JSON(fiber.Map{
		"name": user.Name,
	})
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var register user_models.Register

	if err := c.BodyParser(&register); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if !isValidEmail(register.Email) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid email address format. Please use ___@___.___"})
	}

	if !isValidPassword((register.Password)) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid password format. Must be 10+ characters long, have at least 1 Uppercase, at least 1 Lowercase, and at least 1 number."})
	}

	email, err := h.userService.RegisterUser(register.Email, register.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Registeration Error: Please contact the owner of the site."})
	}

	user := user_models.User{
		Name: register.Name,
	}

	if err := h.userService.CreateGraphUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Registeration Error: Please contact the owner of the site."})
	}

	return c.Status(201).JSON(fiber.Map{
		"username": user.Name,
		"email":    email,
	})
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var login user_models.Login

	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if !isValidEmail(login.Email) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid email address format. Please use ___@___.___"})
	}

	if !isValidPassword((login.Password)) {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid password format. Must be 10+ characters long, have at least 1 Uppercase, at least 1 Lowercase, and at least 1 number."})
	}

	token, err := h.userService.SigninUser(login.Email, login.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
	}

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}

func isValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidPassword(password string) bool {
	if len(password) < 10 {
		return false
	}

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)

	return hasLower && hasUpper && hasDigit
}

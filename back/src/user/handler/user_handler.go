package handler

import (
	"back/src/user/models"
	"back/src/user/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	user, err := h.userService.GetUser()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not get user"})
	}
	return c.JSON(user)
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var register models.Register

	// Parse the request body into the user struct
	if err := c.BodyParser(&register); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Attempt to register the user with Supabase
	token, err := h.userService.RegisterUser(register.Email, register.Password)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "could not create Supabase user: " + err.Error()})
	}

	user := models.User{
		Name: register.Name,
	}

	// If Supabase user creation is successful, proceed to create the graph user
	if err := h.userService.CreateGraphUser(&user); err != nil {
		// Optional: You may want to handle the case where Graph user creation fails
		return c.Status(500).JSON(fiber.Map{"error": "could not create graph user: " + err.Error()})
	}

	// Return the created user along with the JWT token
	return c.Status(201).JSON(fiber.Map{
		"user":  user,
		"token": token, // Include the JWT token if needed
	})
}

// Login handler to authenticate the user
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var login models.Login

	// Parse the request body into the user struct
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Attempt to sign in the user with Supabase
	token, err := h.userService.SigninUser(login.Email, login.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid credentials: " + err.Error()})
	}

	// Return the authenticated user along with the JWT token
	return c.Status(200).JSON(fiber.Map{
		"token": token, // Include the JWT token for authenticated access
	})
}

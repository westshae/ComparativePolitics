package handlers

import (
	"back/src/models"
	"back/src/services"

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

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := h.userService.CreateUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "could not create user"})
	}

	return c.Status(201).JSON(user)
}

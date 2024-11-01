package question_handler

import (
	"back/src/question/question_models"
	"back/src/question/question_services"

	"github.com/gofiber/fiber/v2"
)

type QuestionHandler struct {
	questionService *question_services.QuestionService
}

func NewUserHandler(userService *question_services.QuestionService) *QuestionHandler {
	return &QuestionHandler{
		questionService: userService,
	}
}

// func (h *QuestionHandler) GetUser(c *fiber.Ctx) error {
// 	name := c.Query("name")
// 	if name == "" {
// 		return c.Status(400).JSON(fiber.Map{"error": "name parameter is required"})
// 	}

// 	user, err := h.questionService.GetGraphUser(name)
// 	if err != nil {
// 		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
// 	}

// 	return c.Status(200).JSON(fiber.Map{
// 		"name": user.Name,
// 	})
// }

func (h *QuestionHandler) CreateSide(c *fiber.Ctx) error {
	var sideRequest question_models.SideRequest

	if err := c.BodyParser(&sideRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	id, err := h.questionService.CreateSide(sideRequest.Statement)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create side"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id": id,
	})
}

func (h *QuestionHandler) CreateQuestion(c *fiber.Ctx) error {
	var questionRequest question_models.QuestionRequest

	if err := c.BodyParser(&questionRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	id, err := h.questionService.CreateQuestion(questionRequest.Combiner, questionRequest.LeftSideId, questionRequest.RightSideId)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create question"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id": id,
	})
}

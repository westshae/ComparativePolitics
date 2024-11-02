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

func (h *QuestionHandler) GetAllSides(c *fiber.Ctx) error {
	user, err := h.questionService.GetAllSides()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Failed to get all sides"})
	}

	return c.Status(200).JSON(fiber.Map{
		"sides": user,
	})
}

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

func (h *QuestionHandler) GetQuestion(c *fiber.Ctx) error {
	leftId, leftStatement, rightId, rightStatement, err := h.questionService.GetQuestion()

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to get question"})
	}

	return c.Status(201).JSON(fiber.Map{
		"leftId":         leftId,
		"leftStatement":  leftStatement,
		"rightId":        rightId,
		"rightStatement": rightStatement,
	})
}

func (h *QuestionHandler) CreateAnswer(c *fiber.Ctx) error {
	var answerRequest question_models.AnswerRequest

	if err := c.BodyParser(&answerRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	id, err := h.questionService.CreateAnswer(answerRequest.Username, answerRequest.Preferred, answerRequest.Unpreferred)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Unable to create answer"})
	}

	return c.Status(201).JSON(fiber.Map{
		"id": id,
	})
}

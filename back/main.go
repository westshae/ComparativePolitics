package main

import (
	"github.com/gofiber/fiber/v2"
)

// User struct represents a simple data model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		user := User{
			ID:   1,
			Name: "John Doe",
			Age:  30,
		}
		return c.JSON(user)
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		// Define a variable to hold the incoming user data
		var user User

		// Parse the JSON body
		if err := c.BodyParser(&user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
		}

		// You can process the user data here (e.g., save to a database)

		// Respond with the created user data
		return c.Status(201).JSON(user)
	})

	// Start server on port 3000
	app.Listen(":5000")
}

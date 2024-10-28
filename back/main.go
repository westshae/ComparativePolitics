package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// User struct represents a simple data model
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var driver neo4j.Driver

func initNeo4j() {
	var err error
	// Connect to the Neo4j database
	driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password123", ""))
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Initialize Neo4j connection
	initNeo4j()
	defer driver.Close()

	// Create a new Fiber app
	app := fiber.New()

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!")
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

		// Save the user data to Neo4j
		session := driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
		defer session.Close()

		_, err := session.Run("CREATE (u:User {id: $id, name: $name, age: $age})", map[string]interface{}{
			"id":   user.ID,
			"name": user.Name,
			"age":  user.Age,
		})
		if err != nil {
			fmt.Println("Error could not create user:", err) // Print error to console

			return c.Status(500).JSON(fiber.Map{"error": "could not create user"})
		}

		// Respond with the created user data
		return c.Status(201).JSON(user)
	})

	// Start server on port 5000
	app.Listen(":5000")
}

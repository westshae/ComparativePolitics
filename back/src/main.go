package main

import (
	"log"

	"back/src/handlers"
	"back/src/services"

	"github.com/gofiber/fiber/v2"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	// Initialize Neo4j connection
	driver, err := initNeo4j()
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	// Initialize services
	userService := services.NewUserService(driver)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Create a new Fiber app
	app := fiber.New()

	// Routes
	app.Get("/", homeHandler)
	app.Get("/user", userHandler.GetUser)
	app.Post("/user", userHandler.CreateUser)

	// Start server on port 5000
	app.Listen(":5000")
}

func initNeo4j() (neo4j.Driver, error) {
	return neo4j.NewDriver(
		"bolt://localhost:7687",
		neo4j.BasicAuth("neo4j", "password123", ""),
	)
}

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!")
}

package main

import (
	"log"
	"os"

	"back/src/user/handler"
	"back/src/user/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
	userHandler := handler.NewUserHandler(userService)

	// Create a new Fiber app
	app := fiber.New()

	// Routes
	app.Get("/", homeHandler)
	app.Get("/user", userHandler.GetUser)
	app.Post("/user", userHandler.CreateUser)
	app.Post("/login", userHandler.Login)

	// Start server on port 5000
	app.Listen(":5000")
}

func initNeo4j() (neo4j.Driver, error) {
	err := godotenv.Load("../")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")

	return neo4j.NewDriver(
		dbHost,
		neo4j.BasicAuth(dbUser, dbPassword, ""),
	)
}

func homeHandler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!!")
}

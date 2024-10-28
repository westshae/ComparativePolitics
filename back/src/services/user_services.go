package services

import (
	"fmt"

	"back/src/models"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type UserService struct {
	driver neo4j.Driver
}

func NewUserService(driver neo4j.Driver) *UserService {
	return &UserService{
		driver: driver,
	}
}

func (s *UserService) GetUser() (*models.User, error) {
	// This is just a sample implementation
	return &models.User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
	}, nil
}

func (s *UserService) CreateUser(user *models.User) error {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	_, err := session.Run(
		"CREATE (u:User {id: $id, name: $name, age: $age})",
		map[string]interface{}{
			"id":   user.ID,
			"name": user.Name,
			"age":  user.Age,
		},
	)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

package services

import (
	"fmt"
	"os"

	"back/src/user/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
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
		Name: "John Doe",
	}, nil
}

func (s *UserService) CreateGraphUser(user *models.User) error {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	_, err := session.Run(
		"CREATE (u:User {name: $name})",
		map[string]interface{}{
			"name": user.Name,
		},
	)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (s *UserService) SigninUser(email string, password string) (string, error) {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ANON_KEY"), &supabase.ClientOptions{})
	if err != nil {
		return "", fmt.Errorf("cannot initialize client: %w", err)
	}

	session, err := client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		return "", fmt.Errorf("signin failed: %w", err)
	}

	return session.AccessToken, nil
}

// RegisterUser registers a new user and returns the JWT token
func (s *UserService) RegisterUser(email string, password string) (string, error) {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ANON_KEY"), &supabase.ClientOptions{})
	if err != nil {
		return "", fmt.Errorf("cannot initialize client: %w", err)
	}

	signupRequest := types.SignupRequest{
		Email:    email,
		Password: password,
	}

	// Call the SignUp method to register a new user
	user, err := client.Auth.Signup(signupRequest)
	if err != nil {
		return "", fmt.Errorf("registration failed: %w", err)
	}
	fmt.Printf("New user registered in: %s", user.Email)

	return user.Email, nil
}

// ValidateJWT checks if the JWT is valid
func ValidateJWT(tokenString string) (bool, error) {
	// You need to replace `your-jwt-signing-key` with your actual signing key
	secretKey := []byte(os.Getenv("JWT_KEY"))

	// Parse the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false, fmt.Errorf("invalid token: %w", err)
	}

	// Check if the token is valid
	return token.Valid, nil
}

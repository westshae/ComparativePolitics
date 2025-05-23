package user_services

import (
	"fmt"
	"os"

	"back/src/user/user_models"

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

func (s *UserService) GetGraphUserViaName(name string) (*user_models.User, error) {
	// Start a new session
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	// Run the MATCH query to retrieve the user by name
	result, err := session.Run(
		"MATCH (u:User {name: $name}) RETURN u",
		map[string]interface{}{"name": name},
	)
	if err != nil {
		return nil, fmt.Errorf("could not find user: %w", err)
	}

	if result.Next() {
		record := result.Record()
		node, ok := record.Get("u")
		if !ok {
			return nil, fmt.Errorf("user not found")
		}

		userNode := node.(neo4j.Node)
		userEmail := userNode.Props["email"].(string)
		userName := userNode.Props["name"].(string)

		return &user_models.User{Name: userName, Email: userEmail}, nil
	}

	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	return nil, fmt.Errorf("user not found")
}

func (s *UserService) GetGraphUserViaEmail(email string) (*user_models.User, error) {
	// Start a new session
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	// Run the MATCH query to retrieve the user by name
	result, err := session.Run(
		"MATCH (u:User {email: $email}) RETURN u",
		map[string]interface{}{"email": email},
	)
	if err != nil {
		return nil, fmt.Errorf("could not find user: %w", err)
	}

	if result.Next() {
		record := result.Record()
		node, ok := record.Get("u")
		if !ok {
			return nil, fmt.Errorf("user not found")
		}

		userNode := node.(neo4j.Node)
		userEmail := userNode.Props["email"].(string)
		userName := userNode.Props["name"].(string)

		return &user_models.User{Name: userName, Email: userEmail}, nil
	}

	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("error retrieving user: %w", err)
	}

	return nil, fmt.Errorf("user not found")
}

func (s *UserService) CreateGraphUser(user *user_models.User) error {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	_, err := session.Run(
		"CREATE (u:User {name: $name, email: $email})",
		map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		},
	)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (s *UserService) SigninUser(email string, password string) (string, *user_models.User, error) {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ANON_KEY"), &supabase.ClientOptions{})
	if err != nil {
		return "", nil, fmt.Errorf("cannot initialize client: %w", err)
	}

	session, err := client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		return "", nil, fmt.Errorf("signin failed: %w", err)
	}

	user, err := s.GetGraphUserViaEmail(email)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get user from db: %w", err)
	}

	return session.AccessToken, user, nil
}

func (s *UserService) RegisterUser(email string, password string) (string, error) {
	client, err := supabase.NewClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ANON_KEY"), &supabase.ClientOptions{})
	if err != nil {
		return "", fmt.Errorf("cannot initialize client: %w", err)
	}

	signupRequest := types.SignupRequest{
		Email:    email,
		Password: password,
	}

	user, err := client.Auth.Signup(signupRequest)
	if err != nil {
		return "", fmt.Errorf("registration failed: %w", err)
	}
	fmt.Printf("New user registered in: %s", user.Email)

	return user.Email, nil
}

func ValidateJWT(tokenString string) (bool, error) {
	secretKey := []byte(os.Getenv("JWT_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false, fmt.Errorf("invalid token: %w", err)
	}

	return token.Valid, nil
}

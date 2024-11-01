package question_services

import (
	"fmt"
	"strconv"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type QuestionService struct {
	driver neo4j.Driver
}

func NewQuestionService(driver neo4j.Driver) *QuestionService {
	return &QuestionService{
		driver: driver,
	}
}

func (s *QuestionService) CreateSide(statement string) (int64, error) {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	// Run the query and capture the result
	result, err := session.Run(
		"CREATE (s:Side {statement: $statement}) RETURN id(s) AS sideID",
		map[string]interface{}{
			"statement": statement,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("could not create side: %w", err)
	}

	// Retrieve the generated ID from the query result
	if result.Next() {
		sideID := result.Record().GetByIndex(0).(int64)
		return sideID, nil
	}

	return 0, fmt.Errorf("could not retrieve side ID")
}

func (s *QuestionService) CreateQuestion(combiner string, leftSide string, rightSide string) (int64, error) {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	leftSideID, err := strconv.ParseInt(leftSide, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid leftSide ID: %w", err)
	}

	rightSideID, err := strconv.ParseInt(rightSide, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid rightSide ID: %w", err)
	}

	result, err := session.Run(
		`CREATE (q:Question {combiner: $combiner}) 
		 WITH q
		 MATCH (s1:Side) WHERE id(s1) = $leftSideID
		 MATCH (s2:Side) WHERE id(s2) = $rightSideID
		 CREATE (q)-[:LEFT_SIDE]->(s1) 
		 CREATE (q)-[:RIGHT_SIDE]->(s2)`,
		map[string]interface{}{
			"combiner":    combiner,
			"leftSideID":  leftSideID,
			"rightSideID": rightSideID,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("could not create question: %w", err)
	}

	if result.Next() {
		sideID := result.Record().GetByIndex(0).(int64)
		return sideID, nil
	}

	return 0, fmt.Errorf("could not retrieve question ID")
}

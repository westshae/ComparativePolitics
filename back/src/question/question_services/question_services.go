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
		 CREATE (q)-[:RIGHT_SIDE]->(s2)
		 RETURN id(q) AS questionId`,
		map[string]interface{}{
			"combiner":    combiner,
			"leftSideID":  leftSideID,
			"rightSideID": rightSideID,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("could not create question: %w", err)
	}
	fmt.Printf("%s", result.Record())
	if result.Next() {
		sideID := result.Record().GetByIndex(0).(int64)
		return sideID, nil
	}

	fmt.Println("Error below??")
	return 0, fmt.Errorf("could not retrieve question ID")
}

func (s *QuestionService) GetAllQuestions() ([]map[string]interface{}, error) {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	// Run the query to get questions with their associated sides
	result, err := session.Run(
		`MATCH (q:Question)-[:LEFT_SIDE]->(left:Side), 
		       (q)-[:RIGHT_SIDE]->(right:Side) 
		 RETURN id(q) AS questionID, q.combiner AS combiner,
		        id(left) AS leftSideID, left.statement AS leftStatement,
		        id(right) AS rightSideID, right.statement AS rightStatement`,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve questions: %w", err)
	}

	// Collect results into a slice
	var questions []map[string]interface{}
	for result.Next() {
		record := result.Record()
		question := map[string]interface{}{
			"questionID":     record.GetByIndex(0),
			"combiner":       record.GetByIndex(1),
			"leftSideID":     record.GetByIndex(2),
			"leftStatement":  record.GetByIndex(3),
			"rightSideID":    record.GetByIndex(4),
			"rightStatement": record.GetByIndex(5),
		}
		questions = append(questions, question)
	}

	// Check for any errors after iterating
	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through results: %w", err)
	}

	return questions, nil
}

func (s *QuestionService) GetAllSides() ([]map[string]interface{}, error) {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	// Run the query to get all sides
	result, err := session.Run(
		`MATCH (s:Side) 
		 RETURN id(s) AS sideID, s.statement AS statement`,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve sides: %w", err)
	}

	// Collect results into a slice
	var sides []map[string]interface{}
	for result.Next() {
		record := result.Record()
		side := map[string]interface{}{
			"sideID":    record.GetByIndex(0),
			"statement": record.GetByIndex(1),
		}
		sides = append(sides, side)
	}

	// Check for any errors after iterating
	if err = result.Err(); err != nil {
		return nil, fmt.Errorf("error iterating through results: %w", err)
	}

	return sides, nil
}

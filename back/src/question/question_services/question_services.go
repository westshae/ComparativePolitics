package question_services

import (
	"fmt"

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

func (s *QuestionService) GetQuestion() (int64, string, int64, string, error) {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	result, err := session.Run(
		`MATCH (s:Side) 
		 WITH s ORDER BY rand() 
		 RETURN id(s) AS sideId, s.statement AS statement LIMIT 2`,
		nil,
	)
	if err != nil {
		return 0, "", 0, "", fmt.Errorf("could not retrieve random sides: %w", err)
	}

	// Collect the two side IDs and statements
	var sideIDs []int64
	var statements []string
	for result.Next() {
		sideID := result.Record().GetByIndex(0).(int64)
		statement := result.Record().GetByIndex(1).(string)
		sideIDs = append(sideIDs, sideID)
		statements = append(statements, statement)
	}

	// Ensure we have exactly two sides
	if len(sideIDs) != 2 || len(statements) != 2 {
		return 0, "", 0, "", fmt.Errorf("could not retrieve two unique sides")
	}

	return sideIDs[0], statements[0], sideIDs[1], statements[1], nil
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

func (s *QuestionService) CreateAnswer(userName string, preferredSideID string, unpreferredSideID string) (int64, error) {
	session := s.driver.NewSession(neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close()

	result, err := session.Run(
		`MATCH (u:User {name: $userName})
		 MATCH (preferred:Side) WHERE id(preferred) = toInteger($preferredSideID)
		 MATCH (unpreferred:Side) WHERE id(unpreferred) = toInteger($unpreferredSideID)
		 CREATE (a:Answer)-[:ANSWERED]->(u)
		 CREATE (a)-[:PREFERRED]->(preferred)
		 CREATE (a)-[:UNPREFERRED]->(unpreferred)
		 RETURN id(a) AS answerID`,
		map[string]interface{}{
			"userName":          userName,
			"preferredSideID":   preferredSideID,
			"unpreferredSideID": unpreferredSideID,
		},
	)
	if err != nil {
		return 0, fmt.Errorf("could not create answer: %w", err)
	}

	if result.Next() {
		answerID := result.Record().GetByIndex(0).(int64)
		return answerID, nil
	}

	return 0, fmt.Errorf("could not retrieve answer ID")
}

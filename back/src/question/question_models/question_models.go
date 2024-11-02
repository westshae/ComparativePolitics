package question_models

type SideRequest struct {
	Statement string `json:"statement"`
}

type AnswerRequest struct {
	Username    string `json:"username"`
	QuestionId  string `json:"questionid"`
	Preferred   string `json:"preferred"`
	Unpreferred string `json:"unpreferred"`
}

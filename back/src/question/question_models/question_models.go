package question_models

type Answer struct {
	Question Question
}

type Question struct {
	LeftSide  Side
	RightSide Side
}

type QuestionRequest struct {
	LeftSideId  string `json:"leftSideId"`
	RightSideId string `json:"rightSideId"`
	Combiner    string `json:"combiner"`
}

type Side struct {
	Statement string
}

type SideRequest struct {
	Statement string `json:"statement"`
}

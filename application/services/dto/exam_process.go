package dto

type ExamProcess struct {
	ExamID    string          `json:"exam_id"`
	StudentID string          `json:"student_id"`
	Questions []*QuestionInfo `json:"questions"`
}

type QuestionInfo struct {
	QuestionID string        `json:"question_id"`
	Actions    []*ActionInfo `json:"actions"`
	Score      uint          `json:"score"`
	ScoredBy   string        `json:"scored_by"`
	ScoredTime int64         `json:"scored_time"`
}

type ActionInfo struct {
	ActionID   string `json:"action_id"`
	Answer     string `json:"answer"`
	ActionTime int64  `json:"action_time"`
}

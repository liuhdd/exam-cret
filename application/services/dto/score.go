package dto

type Score struct {
	QuestionID string `json:"question_id"`
	ActionID   string `json:"action_id"`
	Score      uint   `json:"score"`
	ScoredBy   string `json:"scored_by"`
	ScoredTime int64  `json:"scored_time"`
}

type Question struct {
	ExamID     string `json:"exam_id"`
	StudentID  string `json:"student_id"`
	QuestionID string `json:"question_id"`
	Score      uint   `json:"score"`
}

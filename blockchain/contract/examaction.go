package main

const (
	TEXT uint = iota + 1

	MULTI_CHOICE

	FILE_SUMMARY
)

type ExamAction struct {
	ObjectType string `json:"object_type"`
	ActionID   string `json:"action_id"`
	ExamID     string `json:"exam_id"`
	StudentID  string `json:"student_id"`
	ActionType uint   `json:"action_type"`
	ActionTime int64  `json:"action_time"`
	QuestionID string `json:"question_id"`
	Answer     string `json:"answer"`
}

type MarkAction struct {
	ObjectType string `json:"object_type"`
	ActionID   string `json:"action_id"`
	ExamID     string `json:"exam_id"`
	StudentID  string `json:"student_id"`
	QuestionID string `json:"question_id"`
	Score	  uint   `json:"score"`
	Scorer    string `json:"scorer"`
	ScoredTime int64  `json:"scored_time"`
}

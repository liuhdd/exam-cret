package main

const (
	TEXT uint = iota + 1

	MULTI_CHOICE

	FILE_SUMMARY
)

type ExamAction struct {
	ActionID   string `json:"action_id"`
	ExamID     string `json:"exam_id"`
	StudentID  string `json:"student_id"`
	ActionType uint   `json:"action_type"`
	ActionTime int64  `json:"action_time"`
	QuestionID string `json:"question_id"`
	Answer     string `json:"answer"`
}

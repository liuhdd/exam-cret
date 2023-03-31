package models

type ExamAction struct {
	ObjectType string `json:"object_type"`
	ActionID   string `json:"action_id" gorm:"primary_key"`
	ExamID     string `json:"exam_id" gorm:"not null"`
	StudentID  string `json:"student_id" gorm:"not null"`
	ActionType uint   `json:"action_type" gorm:"not null"`
	ActionTime int64  `json:"action_time" gorm:"not null"`
	QuestionID string `json:"question_id" gorm:"not null"`
	Answer     string `json:"answer"`
}

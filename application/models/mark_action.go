package models

type MarkAction struct {
	ObjectType string `json:"object_type"`
	ActionID   string `json:"action_id" gorm:"primary_key"`
	ExamID     string `json:"exam_id" gorm:"not null"`
	StudentID  string `json:"student_id" gorm:"not null"`
	QuestionID string `json:"question_id" gorm:"not null"`
	Score	  uint   `json:"score" gorm:"not null"`
	Scorer    string `json:"scorer" gorm:"not null"`
	ScoredTime int64  `json:"scored_time" gorm:"not null"`
}
package models

type MarkAction struct {
	ObjectType string `json:"object_type"gorm:"default:'mark_action'"`
	ActionID   string `json:"action_id" gorm:"primary_key"`
	ExamID     string `json:"exam_id" gorm:"not null"`
	StudentID  string `json:"student_id" gorm:"not null"`
	QuestionID string `json:"question_id" gorm:"not null"`
	Score      uint   `json:"score" gorm:"not null"`
	Scorer     string `json:"scorer" gorm:"not null"`
	ActionTime int64  `json:"action_time" gorm:"not null"`
}

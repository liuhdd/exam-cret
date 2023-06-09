package models

type Mark struct {
	ExamID     string `json:"exam_id" gorm:"not null;primaryKey" form:"exam_id"`
	StudentID  string `json:"student_id" gorm:"not null;primaryKey" form:"student_id"`
	QuestionID string `json:"question_id" gorm:"not null;primaryKey" form:"question_id"`
	Answer     string `json:"answer" form:"answer"`
	Score      uint   `json:"score" form:"score;default:0"`
}

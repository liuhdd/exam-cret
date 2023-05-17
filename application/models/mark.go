package models

type Mark struct {
	ExamID     string `json:"exam_id" gorm:"not null;primary_key" form:"exam_id"`
	StudentID  string `json:"student_id" gorm:"not null;primary_key" form:"student_id"`
	QuestionID string `json:"question_id" gorm:"not null;primary_key" form:"question_id"`
	Answer     string `json:"answer" form:"answer"`
	Score      uint   `json:"score" form:"score"`
}

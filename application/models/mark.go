package models

type Mark struct {
	ExamID    string `json:"exam_id" gorm:"not null" form:"exam_id"`
	StudentID string `json:"student_id" gorm:"not null" form:"student_id"`
	QuestionID string `json:"question_id" gorm:"not null" form:"question_id"`
	Answer string `json:"answer" gorm:"not null" form:"answer"`
	Score uint `json:"score" gorm:"not null" form:"score"`
}
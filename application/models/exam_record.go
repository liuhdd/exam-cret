package models

type ExamRecord struct {
	ExamID    string `json:"exam_id" gorm:"not null;primaryKey'" form:"exam_id"`
	StudentID string `json:"student_id" gorm:"not null;primaryKey" form:"student_id"`
	Grade     int    `json:"grade" gorm:"not null" form:"grade"`
}

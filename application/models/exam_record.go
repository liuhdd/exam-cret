package models

type ExamRecord struct {
	ExamRecordID   string `json:"exam_record_id" gorm:"primary_key" form:"exam_record_id"`
	ExamID string `json:"exam_id" gorm:"not null" form:"exam_id"`
	StudentID string `json:"student_id" gorm:"not null" form:"student_id"`
}
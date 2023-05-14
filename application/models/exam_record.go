package models

type ExamRecord struct {
	ExamID    string `json:"exam_id" gorm:"not null" form:"exam_id"`
	StudentID string `json:"student_id" gorm:"not null" form:"student_id"`
	Score     uint    `json:"score" gorm:"not null" form:"score"`
	Grade    float32 `json:"grade" gorm:"not null" form:"grade"`
	UpdateTime int64  `json:"update_time" gorm:"not null" form:"update_time"`
}

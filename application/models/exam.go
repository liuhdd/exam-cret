package models

type Exam struct {
	ExamID    string `json:"exam_id" gorm:"primary_key" form:"exam_id"`
	ExamName  string `json:"exam_name" gorm:"not null" form:"exam_name"`
	BeginTime int64  `json:"begin_time" gorm:"not null" form:"begin_time"`
	EndTime   int64  `json:"end_time" gorm:"not null" form:"end_time"`
	ExamDesc  string `json:"exam_desc"  form:"exam_desc"`
}

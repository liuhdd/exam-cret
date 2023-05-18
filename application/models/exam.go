package models

type Exam struct {
	ExamID        string `json:"exam_id" gorm:"primaryKey" form:"exam_id"`
	ExamName      string `json:"exam_name" gorm:"not null" form:"exam_name"`
	Paper         string `json:"paper" gorm:"not null"`
	Place         string `json:"place" gorm:"not null"`
	BeginTime     int64  `json:"begin_time" gorm:"not null" form:"begin_time"`
	EndTime       int64  `json:"end_time" gorm:"not null" form:"end_time"`
	State         int    `json:"state" gorm:"default:0"`
	ExamServer    string `json:"exam_server"`
	ExamServerKey string `json:"exam_server_key"`
	ExamDesc      string `json:"exam_desc"  form:"exam_desc"`
}

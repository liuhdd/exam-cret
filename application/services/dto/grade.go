package dto

type Grade struct {
	ExamID      string  `json:"exam_id"`
	ExamName    string  `json:"exam_name"`
	StudentID   string  `json:"student_id"`
	StudentName string  `json:"student_name"`
	BeginTime   int64   `json:"begin_time"`
	EndTime     int64   `json:"end_time"`
	Grade       float32 `json:"grade"`
}

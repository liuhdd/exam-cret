package dto

type ExamResult struct {
	ExamID    string            `json:"exam_id"`
	ExamName  string            `json:"exam_name"`
	BeginTime int64             `json:"begin_time"`
	EndTime   int64             `json:"end_time"`
	StudentID string            `json:"student_id"`
	Place     string            `json:"place"`
	Paper     string            `json:"paper"`
	Grade     int               `json:"grade"`
	Questions []*QuestionResult `json:"questions"`
}

type QuestionResult struct {
	QuestionID   string `json:"question_id"`
	Content      string `json:"content"`
	QuestionType string `json:"question_type"`
	Options      string `json:"options"`
	Answer       string `json:"answer"`
	Score        uint   `json:"score"`
}

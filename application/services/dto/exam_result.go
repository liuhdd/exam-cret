package dto

type ExamResult struct {
	ExamID    string `json:"exam_id"`
	StudentID string `json:"student_id"`
	Questions []*QuestionResult `json:"questions"`
}

type QuestionResult struct {
	QuestionID string `json:"question_id"`
	Answer     string `json:"answer"`
}
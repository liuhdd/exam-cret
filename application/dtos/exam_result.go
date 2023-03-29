package dtos

type ExamResult struct {
	ExamID    string `json:"exam_id"`
	StudentID string `json:"student_id"`
	Questions []struct {
		QuestionID string `json:"question_id"`
		Answer     string `json:"answer"`
	} `json:"questions"`
}


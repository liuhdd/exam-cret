package dto

type QuestionActions struct {
	ExamID    string `json:"exam_id"`
	StudentID string `json:"student_id"`
	QuestionID string `json:"question_id"`
	Actions []struct {
		ActionID   string `json:"action_id"`
		Answer	 string `json:"answer"`
		ActionTime int64  `json:"action_time"`
	} `json:"actions"`
}
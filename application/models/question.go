package models

type Question struct {
	QuestionID string `json:"question_id" gorm:"primary_key" form:"question_id"`
	QuestionType string `json:"question_type" gorm:"not null" form:"question_type"`
	Content string `json:"_content" gorm:"not null" form:"content"`
	Solution string `json:"solution" gorm:"not null" form:"solution"`
	Score int `json:"_score" gorm:"not null" form:"score"`
}
package models

type Student struct {
	User
	StudentID   string `json:"student_id" gorm:"primary_key" form:"student_id"`
	Name string `json:"name" gorm:"not null" form:"name"`
}
package models

type Student struct {
	User
	StudentID string `json:"student_id" gorm:"primary_key" form:"student_id"`

	name string `json:"name" gorm:"not null" form:"name"`
}

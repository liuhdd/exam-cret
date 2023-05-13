package models

type Teacher struct {
	User
	TeacherID string `json:"teacher_id" gorm:"primary_key" form:"teacher_id"`
	Name string `json:"name" gorm:"not null" form:"name"`
	Gender string `josn:"gender" gorm:"not null" form:"gender"`
	Email string `json:"email" form:"email"`
}
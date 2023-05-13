package models

type Teacher struct {
	User
	TeacherID string `json:"teacher_id" gorm:"not null" form:"teacher_id"`
	Name string `json:"name" gorm:"not null" form:"name"`
	Gender string `josn:"gender" form:"gender"`
	Email string `json:"email" form:"email"`
}
package models

type Student struct {
	User
	StudentID string `json:"student_id" gorm:"not null;unique" form:"student_id"`
	Name string `json:"name" gorm:"not null" form:"name"`
	Gender string `json:"gender" form:"email"`
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
}

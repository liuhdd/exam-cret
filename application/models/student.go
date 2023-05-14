package models

type Student struct {
	User
	StudentID string `json:"student_id" gorm:"not null;unique" form:"student_id" redis:"student_id"`
	Name      string `json:"name" gorm:"not null" form:"name" redis:"name"`
	Gender    string `json:"gender" form:"email" redis:"gender"`
	Email     string `json:"email" form:"email" redis:"email"`
	Phone     string `json:"phone" form:"phone" redis:"phone"`
}

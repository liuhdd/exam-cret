package models

type Teacher struct {
	User
	TeacherID string `json:"teacher_id" gorm:"not null" form:"teacher_id" redis:"teacher_id"`
	Name      string `json:"name" gorm:"not null" form:"name" redis:"name"`
	Gender    string `josn:"gender" form:"gender" redis:"gender"`
	Email     string `json:"email" form:"email" redis:"email"`
	Phone     string `json:"phone" form:"phone" redis:"phone"`
}

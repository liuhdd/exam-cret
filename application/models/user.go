package models

type User struct {
	UserID   string `json:"user_id" gorm:"primary_key" form:"user_id" redis:"user_id"`
	Username string `json:"username" gorm:"unique;not null" form:"username" redis:"username"`
	Password string `json:"password" gorm:"not null" form:"password" redis:"password"`
	Role     string `json:"role" form:"role" redis:"role"`
}

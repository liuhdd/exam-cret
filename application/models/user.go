package models

type User struct {
	UserID   string `json:"user_id" gorm:"primary_key" form:"user_id"`
	Username string `json:"username" gorm:"unique;not null" form:"username"`
	Password string `json:"password" gorm:"not null" form:"username"`
}

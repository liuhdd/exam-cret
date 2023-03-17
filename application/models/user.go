package models

type User struct {
	UserID   string
	Username string `json:"username"gorm:"unique;not null"form:"username"'`
	Password string `json:"password"gorm:"not null"form:"username"'`
}

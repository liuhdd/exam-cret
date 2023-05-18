package models

type User struct {
	UserID        int64  `json:"user_id" gorm:"primaryKey;" form:"user_id" redis:"user_id"`
	Username      string `json:"username" gorm:"unique;not null" form:"username" redis:"username"`
	Password      string `json:"password" gorm:"not null" form:"password" redis:"password"`
	Role          string `json:"role" form:"role" redis:"role"`
	UpdateTime    int64  `json:"update_time" gorm:"not null" form:"update_time"`
	LastLoginTime int64  `json:"last_login_time" gorm:"not null" form:"last_login_time"`
}

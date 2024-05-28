package domain

import (
	"time"
)

type User struct {
	UserID      string `gorm:"column:id;"`
	Name        string `gorm:"column:name"`
	Username    string `gorm:"column:username"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	Phonenumber string `gorm:"column:phonenumber"`
	CreatedAt   time.Time
}

func (User) TableName() string {
	return "depublic_users"
}

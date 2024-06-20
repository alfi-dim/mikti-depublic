package domain

import (
	"time"
)

type User struct {
	UserID      string    `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Username    string    `gorm:"column:username" json:"username"`
	Email       string    `gorm:"column:email" json:"email"`
	Password    string    `gorm:"column:password" json:"-"`
	Phonenumber string    `gorm:"column:phonenumber" json:"phonenumber"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"-"`
}

func (User) TableName() string {
	return "depublic_users"
}

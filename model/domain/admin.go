package domain

import "time"

type Admin struct {
	ID        string `gorm:"column:id"`
	Name      string `gorm:"column:name"`
	Username  string `gorm:"column:username"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time
}

func (Admin) TableName() string {
	return "depublic_administrators"
}

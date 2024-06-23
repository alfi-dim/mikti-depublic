package domain

import "time"

type Admin struct {
	ID        string `gorm:"column:id" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Username  string `gorm:"column:username" json:"username"`
	Email     string `gorm:"column:email" json:"email"`
	Password  string `gorm:"column:password" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func (Admin) TableName() string {
	return "depublic_administrators"
}

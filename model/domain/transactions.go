package domain

import "time"

type Transaction struct {
	ID         string  `gorm:"column:id"`
	UserID     string  `gorm:"column:users_id;foreignKey:fk_users_id;references:id"`
	EventID    string  `gorm:"column:events_id;foreignKey:fk_events_id;references:id"`
	Date       string  `gorm:"column:date"`
	Quantity   int     `gorm:"column:quantity"`
	TotalPrice float64 `gorm:"column:total_price"`
	Status     string  `gorm:"column:status"`
	CreatedAt  time.Time
}

func (Transaction) TableName() string {
	return "depublic_transactions"
}

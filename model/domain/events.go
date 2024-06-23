package domain

type Event struct {
	ID          string  `gorm:"column:id"`
	AdminId     string  `gorm:"column:administrators_id;foreignKey:fk_administrators_id;references:id"`
	Name        string  `gorm:"column:name"`
	Address     string  `gorm:"column:address"`
	Date        string  `gorm:"column:date"`
	Price       float64 `gorm:"column:price"`
	Tickets     int     `gorm:"column:tickets"`
	TicketsSold int     `gorm:"column:tickets_sold"`
}

func (Event) TableName() string {
	return "depublic_events"
}

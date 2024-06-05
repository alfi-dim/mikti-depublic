package web

type EventServiceReq struct {
	ID          string  `json:"id"`
	AdminId     string  `validate:"required" json:"administrators_id"`
	Name        string  `validate:"required" json:"name"`
	Address     string  `validate:"required" json:"address"`
	Date        string  `validate:"required" json:"date"`
	Price       float64 `validate:"required" json:"price"`
	Tickets     int     `validate:"required" json:"tickets"`
	TicketsSold int     `validate:"required" json:"tickets_sold"`
}

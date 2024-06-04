package entity

import "mikti-depublic/model/domain"

type EventEntity struct {
	ID          string  `json:"id"`
	AdminId     string  `json:"administrators_id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Date        string  `json:"date"`
	Price       float64 `json:"price"`
	Tickets     int     `json:"tickets"`
	TicketsSold int     `json:"tickets_sold"`
}

func ToEventEntity(id string, administrators_id string, name string, address string, date string, price float64, tickets int, tickets_sold int) EventEntity {
	return EventEntity{
		ID:          id,
		AdminId:     administrators_id,
		Name:        name,
		Address:     address,
		Date:        date,
		Price:       price,
		Tickets:     tickets,
		TicketsSold: tickets_sold,
	}
}

func ToEventListEntity(event []domain.Event) []EventEntity {
	eventData := []EventEntity{}

	for _, events := range event {
		eventData = append(eventData, ToEventEntity(events.ID, events.AdminId, events.Name, events.Address, events.Date, events.Price, events.Tickets, events.TicketsSold))
	}
	return eventData
}

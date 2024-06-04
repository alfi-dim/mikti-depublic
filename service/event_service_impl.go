package service

import (
	"math/rand"
	"mikti-depublic/helper"
	"mikti-depublic/model/domain"
	"mikti-depublic/model/entity"
	"mikti-depublic/model/web"
	"mikti-depublic/repository"
	"time"
)

type EventServiceImpl struct {
	repository repository.EventRepository
}

func NewEventService(repository repository.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		repository: repository,
	}
}

func (service *EventServiceImpl) CreateEvent(request web.EventServiceReq) (map[string]interface{}, error) {
	eventID := generateRandomString(5)
	eventReq := domain.Event{
		ID:          eventID,
		AdminId:     request.AdminId,
		Name:        request.Name,
		Address:     request.Address,
		Date:        request.Date,
		Price:       request.Price,
		Tickets:     request.Tickets,
		TicketsSold: request.TicketsSold,
	}
	saveEvent, errSaveEvent := service.repository.CreateEvent(eventReq)
	if errSaveEvent != nil {
		return nil, errSaveEvent
	}
	return helper.ResponseJson{"id": saveEvent.ID, "name": saveEvent.Name, "address": saveEvent.Address, "date": saveEvent.Date, "price": saveEvent.Price, "tickets": saveEvent.Tickets, "ticket_sold": saveEvent.TicketsSold}, nil
}

func (service *EventServiceImpl) GetEvent(eventId string) (entity.EventEntity, error) {
	getEvent, errGetEvent := service.repository.GetEvent(eventId)

	if errGetEvent != nil {
		return entity.EventEntity{}, errGetEvent
	}
	return entity.ToEventEntity(getEvent.ID, getEvent.AdminId, getEvent.Name, getEvent.Address, getEvent.Date, getEvent.Price, getEvent.Tickets, getEvent.TicketsSold), nil
}

func (service *EventServiceImpl) GetListEvent() ([]entity.EventEntity, error) {
	getEvents, errGetEvents := service.repository.GetListEvent()

	if errGetEvents != nil {
		return []entity.EventEntity{}, errGetEvents
	}
	return entity.ToEventListEntity(getEvents), nil
}

func (service *EventServiceImpl) UpdateEvent(request web.EventServiceReq, pathId string) (map[string]interface{}, error) {
	getEventById, err := service.repository.GetEvent(pathId)

	if err != nil {
		return nil, err
	}

	if request.Name == "" {
		request.Name = getEventById.Name
	}
	if request.Address == "" {
		request.Address = getEventById.Address
	}
	if request.Date == "" {
		request.Date = getEventById.Date
	}
	if request.Price == 0 {
		request.Price = getEventById.Price
	}
	if request.Tickets == 0 {
		request.Tickets = getEventById.Tickets
	}
	if request.TicketsSold == 0 {
		request.TicketsSold = getEventById.TicketsSold
	}

	eventReq := domain.Event{
		ID:          pathId,
		AdminId:     getEventById.AdminId,
		Name:        request.Name,
		Address:     request.Address,
		Date:        request.Date,
		Price:       request.Price,
		Tickets:     request.Tickets,
		TicketsSold: request.TicketsSold,
	}
	updateEvent, err := service.repository.UpdateEvent(eventReq)
	if err != nil {
		return nil, err
	}
	return helper.ResponseJson{"name": updateEvent.Name, "address": updateEvent.Address, "date": updateEvent.Date, "price": updateEvent.Price, "tickets": updateEvent.Tickets, "tickets_sold": updateEvent.TicketsSold}, nil
}

func (service *EventServiceImpl) DeleteEvent(Id string) (entity.EventEntity, error) {
	deleteEvent, errDelete := service.repository.DeleteEvent(Id)
	if errDelete != nil {
		return entity.EventEntity{}, errDelete
	}
	return entity.ToEventEntity(deleteEvent.ID, deleteEvent.AdminId, deleteEvent.Name, deleteEvent.Address, deleteEvent.Date, deleteEvent.Price, deleteEvent.Tickets, deleteEvent.TicketsSold), nil
}

// Fungsi generate random ID untuk event
func generateRandomString(n int) string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	charset := []rune("1234567890")
	letters := make([]rune, n)
	for i := range letters {
		letters[i] = charset[r.Intn(len(charset))]
	}
	return "event-" + string(letters)
}

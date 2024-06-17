package repository

import (
	"errors"
	"mikti-depublic/model/domain"

	"gorm.io/gorm"
)

type EventRepositoryImpl struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db: db}
}

func (repo *EventRepositoryImpl) CreateEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Create(&event).Error

	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (repo *EventRepositoryImpl) GetEvent(Id string) (domain.Event, error) {
	var eventData domain.Event

	err := repo.db.First(&eventData, "id = ?", Id).Error

	if err != nil {
		return domain.Event{}, errors.New("Event not found")
	}

	return eventData, nil
}

func (repo *EventRepositoryImpl) GetListEvent() ([]domain.Event, error) {
	var events []domain.Event

	err := repo.db.Find(&events).Error

	if err != nil {
		return []domain.Event{}, err
	}
	return events, nil
}

// Update Event
func (repo *EventRepositoryImpl) UpdateEvent(event domain.Event) (domain.Event, error) {
	err := repo.db.Model(domain.Event{}).Where("id = ?", event.ID).Updates(event).Error

	if err != nil {
		return event, err
	}
	return event, nil
}

func (repo *EventRepositoryImpl) DeleteEvent(Id string) (domain.Event, error) {
	var eventData domain.Event

	err := repo.db.First(&eventData, "id = ? ", Id).Error

	if err != nil {
		return domain.Event{}, errors.New("Id Event Not found")
	}

	err = repo.db.Delete(&eventData).Error
	if err != nil {
		return domain.Event{}, err
	}

	return eventData, nil
}

func (repo *EventRepositoryImpl) CheckTicketAvailability(eventId string, quantity int) (bool, error) {
	var eventData domain.Event

	err := repo.db.First(&eventData, "id = ?", eventId).Error

	if err != nil {
		return false, errors.New("event not found")
	}

	if eventData.Tickets < quantity {
		return false, errors.New("ticket not available")
	}

	return true, nil
}

func (repo *EventRepositoryImpl) UpdateSoldTicket(eventId string, quantity int) (bool, error) {
	var eventData domain.Event

	err := repo.db.First(&eventData, "id = ?", eventId).Error

	if err != nil {
		return false, errors.New("event not found")
	}

	eventData.TicketsSold += quantity
	eventData.Tickets -= quantity

	err = repo.db.Save(&eventData).Error

	if err != nil {
		return false, errors.New("failed when update event")
	}

	return true, nil
}

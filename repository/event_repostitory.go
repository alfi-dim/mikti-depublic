package repository

import "mikti-depublic/model/domain"

type EventRepository interface {
	CreateEvent(event domain.Event) (domain.Event, error)
	GetEvent(Id string) (domain.Event, error)
	GetListEvent() ([]domain.Event, error)
	UpdateEvent(event domain.Event) (domain.Event, error)
	DeleteEvent(Id string) (domain.Event, error)
}

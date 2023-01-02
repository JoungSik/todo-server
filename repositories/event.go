package repositories

import (
	"todo-server/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.Event) (*models.Event, error)
	Update(event *models.Event) (*models.Event, error)
}

type EventRepositoryImpl struct {
	db *gorm.DB
}

func (EventRepositoryImpl) NewEventEventRepositoryImpl(db *gorm.DB) *EventRepositoryImpl {
	return &EventRepositoryImpl{db}
}

func (eventRepositoryImpl *EventRepositoryImpl) Create(event *models.Event) (*models.Event, error) {
	err := eventRepositoryImpl.db.Create(event).Error
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (eventRepositoryImpl *EventRepositoryImpl) Update(event *models.Event) (*models.Event, error) {
	eventRepositoryImpl.db.Save(&event)
	return event, nil
}

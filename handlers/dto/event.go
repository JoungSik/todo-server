package dto

import (
	"time"
	"todo-server/models"
)

type EventDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartAt     string `json:"start_at"`
}

func (eventDto *EventDto) ToModel() *models.Event {
	startAt, _ := time.Parse("2006-01-02", eventDto.StartAt)
	return &models.Event{Title: eventDto.Title, Description: eventDto.Description, StartAt: startAt}
}

package models

import (
	"time"
)

type Event struct {
	ID          uint `gorm:"primaryKey;auto_increment"`
	Title       string
	Description string
	UserId      uint
	User        User
	StartAt     time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

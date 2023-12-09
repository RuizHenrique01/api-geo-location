package models

import (
	"time"
)

type Point struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"not null"`
	Description string
	Latitude    float64 `gorm:"not null"`
	Longitude   float64 `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

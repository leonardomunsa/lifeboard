package models

import "time"

type Game struct {
	ID        uint   `gorm:"primary_key"`
	Title     string `gorm:"not null"`
	Platform  string `gorm:"not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

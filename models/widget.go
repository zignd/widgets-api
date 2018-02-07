package models

import (
	"time"
)

type Widget struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	ColorID   int       `gorm:"not null"`
	Color     *Color    `gorm:"ForeignKey:ColorID"`
	Price     float64   `gorm:"not null"`
	Melts     bool      `gorm:"not null"`
	Inventory int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

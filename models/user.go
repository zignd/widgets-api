package models

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primary_key"`
	Name      string    `gorm:"not null"`
	Avatar    string    `gorm:"not null"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Admin     bool      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}

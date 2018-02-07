package models

type Color struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
}

package viewmodels

import (
	"time"
)

type Widget struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Color     *Color    `json:"color"`
	Price     float64   `json:"price"`
	Melts     bool      `json:"melts"`
	Inventory int       `json:"inventory"`
	CreatedAt time.Time `json:"createdAt"`
}

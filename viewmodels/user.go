package viewmodels

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Username  string    `json:"username"`
	Admin     bool      `json:"admin"`
	CreatedAt time.Time `json:"createdAt"`
}

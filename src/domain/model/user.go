package model

import "time"

// User model.
type User struct {
	ID        uint
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

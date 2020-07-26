package model

import "time"

// Todo model.
type Todo struct {
	ID          uint
	Title       string
	Description string
	Completed   bool
	DeadlineAt  time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

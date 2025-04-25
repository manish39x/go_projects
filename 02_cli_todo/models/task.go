package models

import "time"

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	DueDate   time.Time `json:"due_date,omitempty"`
	Priority  string    `json:"priority,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
}

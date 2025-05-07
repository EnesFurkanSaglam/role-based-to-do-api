package model

import "time"

type TodoList struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
	Completion float64    `json:"completion"`
	Owner      string     `json:"owner"`
}

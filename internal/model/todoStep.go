package model

import "time"

type TodoStep struct {
	ID        int        `json:"id"`
	ListID    int        `json:"list_id"`
	Content   string     `json:"content"`
	Done      bool       `json:"done"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

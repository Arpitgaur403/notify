package models

import "time"

type Notification struct {
	ID        int        `json:"id"`
	UserId    int        `json:"user_id"`
	Message   string     `json:"message"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	ReadAt    *time.Time `json:"read_at"`
}

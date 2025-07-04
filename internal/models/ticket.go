package models

import "time"

type Ticket struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UserId      int       `json:"user_id"`
	LabelId     int       `json:"label_id"`
	BoardId     int       `json:"board_id"`
}

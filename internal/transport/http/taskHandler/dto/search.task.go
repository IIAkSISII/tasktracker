package dto

import "time"

type SearchTaskResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UserId      int       `json:"user_id"`
	LabelId     int       `json:"label_id"`
	TicketId    int       `json:"ticket_id"`
}

package models

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UserId      int    `json:"user_id"`
	LabelId     int    `json:"label_id"`
	TicketId    int    `json:"ticket_id"`
}

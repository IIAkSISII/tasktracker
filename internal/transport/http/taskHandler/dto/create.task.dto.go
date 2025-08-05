package dto

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	UserId      int    `json:"user_id"`
	LabelId     int    `json:"label_id"`
	TicketId    int    `json:"ticket_id"`
}

type CreateTaskResponse struct {
	Id      int    `json:"id" example:"42"`
	Message string `json:"message" example:"task created"`
}

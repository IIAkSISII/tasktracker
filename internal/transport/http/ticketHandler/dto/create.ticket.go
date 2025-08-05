package dto

type CreateTicketRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	UserId      int    `json:"user_id"`
	LabelId     int    `json:"label_id"`
	BoardId     int    `json:"board_id"`
}

type CreateTicketResponse struct {
	Id      int    `json:"id" example:"42"`
	Message string `json:"message" example:"ticket created"`
}

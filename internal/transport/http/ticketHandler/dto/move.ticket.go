package dto

type MoveTicketRequest struct {
	TicketId   int `json:"ticket_id"`
	NewBoardId int `json:"new_board_id"`
}

type MoveTicketResponse struct {
	Message string `json:"message"`
}

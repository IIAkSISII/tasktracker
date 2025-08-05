package ticketHandler

import (
	"encoding/json"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/ticketHandler/dto"
	"net/http"
)

// @Summary		Переместить карточку на другую доску
// @Description	Перемещаем карточку с указанием id карточки и id новой доски.
// @Description	Параметры передаются в теле запроса в качестве json-объекта.
// @Description	Если карточку не удается создать, возвращаем ошибку.
// @Tags			Tickets
// @Accept			json
// @Produce		json
// @Param			input	body		dto.MoveTicketRequest	true	"ticket_id, new_board_id"
// @Success		201		{object}	dto.MoveTicketResponse	"Карточка успешно перемещена"
// @Failure		400		{string}	string					"Invalid json payload, Name is required или TicketId or BoardId must be positive"
// @Failure		500		{string}	string					"Cannot move ticket: <описание ошибки>"
// @Router			/ticket/move [put]
func (t *ticketHandler) MoveTicket(w http.ResponseWriter, r *http.Request) {
	var req dto.MoveTicketRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if req.TicketId <= 0 || req.NewBoardId <= 0 {
		http.Error(w, "TicketId or BoardId must be positive", http.StatusBadRequest)
		return
	}

	err := t.service.Move(r.Context(), req.TicketId, req.NewBoardId)
	if err != nil {
		http.Error(w, "Cannot move ticket: "+err.Error(), http.StatusInternalServerError)
		t.logger.Error(r.Context(), "Cannot move ticket", logger.Field{Key: "error", Value: err.Error()})
		return
	}

	resp := dto.CreateTicketResponse{
		Message: "Ticket moved",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		t.logger.Error(r.Context(), "Cannot move ticket", logger.Field{Key: "error", Value: err.Error()})
	}
}

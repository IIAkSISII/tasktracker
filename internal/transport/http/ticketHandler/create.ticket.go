package ticketHandler

import (
	"encoding/json"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/ticketHandler/dto"
	"net/http"
)

//	@Summary		Создать новую карточку
//	@Description	Создаем новую карточку с указанием имени, описания, пользователя, метки и доски.
//	@Description	Параметры передаются в теле запроса в качестве json-объекта.
//	@Description	Если карточку не удается создать, возвращаем ошибку.
//	@Tags			Tickets
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dto.CreateTicketRequest		true	"name, description, user_id, label_id, project_id"
//	@Success		201		{object}	dto.CreateTicketResponse	"Карточка успешно создана"
//	@Failure		400		{string}	string						"Invalid json payload, Name is required или LabelId, BoardId or UserId must be positive"
//	@Failure		500		{string}	string						"Cannot create ticket: <описание ошибки>"
//	@Router			/ticket [post]
func (t *ticketHandler) CreateTicket(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTicketRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if req.LabelId <= 0 || req.BoardId <= 0 || req.UserId <= 0 {
		http.Error(w, "LabelId, BoardId or UserId must be positive", http.StatusBadRequest)
		return
	}

	id, err := t.service.Create(r.Context(), req.Name, req.Description, req.UserId, req.LabelId, req.BoardId)
	if err != nil {
		http.Error(w, "Cannot create ticket: "+err.Error(), http.StatusInternalServerError)
		t.logger.Error(r.Context(), "Cannot create ticket", logger.Field{Key: "error", Value: err.Error()})
		return
	}

	resp := dto.CreateTicketResponse{
		Id:      id,
		Message: "Ticket created",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		t.logger.Error(r.Context(), "Cannot create ticket", logger.Field{Key: "error", Value: err.Error()})
	}
}

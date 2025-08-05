package taskHandler

import (
	"encoding/json"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/taskHandler/dto"
	"net/http"
)

// @Summary		Создать новую задачу
// @Description	Создаем новую задачу с указанием имени, описания, пользователя, метки и карточки.
// @Description	Параметры передаются в теле запроса в качестве json-объекта.
// @Description	Если карточку не удается создать, возвращаем ошибку.
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			input	body		dto.CreateTaskRequest	true	"name, description, user_id, label_id, ticket_id"
// @Success		201		{object}	dto.CreateTaskResponse	"Задача успешно создана"
// @Failure		400		{string}	string					"Invalid json payload, Name is required или LabelId, TicketId or UserId must be positive"
// @Failure		500		{string}	string					"Cannot create ticket: <описание ошибки>"
// @Router			/task [post]
func (t *taskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}
	if req.LabelId <= 0 || req.TicketId <= 0 || req.UserId <= 0 {
		http.Error(w, "LabelId, TicketId or UserId must be positive", http.StatusBadRequest)
		return
	}

	id, err := t.service.Create(r.Context(), req.Name, req.Description, req.UserId, req.LabelId, req.TicketId)
	if err != nil {
		http.Error(w, "Cannot create task: "+err.Error(), http.StatusInternalServerError)
		t.logger.Error(r.Context(), "Cannot create task", logger.Field{Key: "error", Value: err.Error()})
		return
	}

	resp := dto.CreateTaskResponse{
		Id:      id,
		Message: "Ticket created",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		t.logger.Error(r.Context(), "Cannot create task", logger.Field{Key: "error", Value: err.Error()})
	}
}

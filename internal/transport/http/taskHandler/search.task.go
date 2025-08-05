package taskHandler

import (
	"encoding/json"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/taskHandler/dto"
	"net/http"
	"strconv"
)

// @Summary		Поиск задач
// @Description	Поиск задач по названию, метке и пользователю
// @Tags			Tasks
// @Accept			json
// @Produce		json
// @Param			name		query		string					false	"Часть названия задачи"
// @Param			label_id	query		int						false	"ID метки"
// @Param			user_id		query		int						false	"ID исполнителя"
// @Success		200			{array}		dto.SearchTaskResponse	"Список найденных задач"
// @Failure		400			{string}	string					"Invalid label_id or user_id"
// @Failure		500			{string}	string					"Failed to search tasks: <описание ошибки>"
// @Router			/task/search [get]
func (t *taskHandler) SearchTask(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	name := query.Get("name")

	labelId := 0
	if v := query.Get("label_id"); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			labelId = id
		} else {
			http.Error(w, "Invalid label_id", http.StatusBadRequest)
			return
		}
	}

	userId := 0
	if v := query.Get("user_id"); v != "" {
		if id, err := strconv.Atoi(v); err == nil {
			userId = id
		} else {
			http.Error(w, "Invalid user_id", http.StatusBadRequest)
			return
		}
	}

	tasks, err := t.service.Search(r.Context(), name, labelId, userId)
	if err != nil {
		http.Error(w, "Failed to search tasks: "+err.Error(), http.StatusInternalServerError)
		t.logger.Error(r.Context(), "Cannot search tasks", logger.Field{Key: "error", Value: err.Error()})
		return
	}

	var resp []dto.SearchTaskResponse
	for _, task := range tasks {
		resp = append(resp, dto.SearchTaskResponse{
			Id:          task.Id,
			Name:        task.Name,
			Description: task.Description,
			CreatedAt:   task.CreatedAt,
			UserId:      task.UserId,
			LabelId:     task.LabelId,
			TicketId:    task.TicketId,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		t.logger.Error(r.Context(), "Failed to search tasks", logger.Field{Key: "error", Value: err.Error()})
	}
}

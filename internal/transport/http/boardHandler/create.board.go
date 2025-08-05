package boardHandler

import (
	"encoding/json"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/boardHandler/dto"
	"net/http"
)

// @Summary		Создать новый доску
// @Description	Создаем новую доску с указанием названия и проекта, к которому она принадлежит.
// @Description	Параметры передаются в теле запроса в качестве json-объекта.
// @Description	Если доску не удается создать, возвращаем ошибку.
// @Tags			Boards
// @Accept			json
// @Produce		json
// @Param			input	body		dto.CreateBoardRequest	true	"name и project_id"
// @Success		201		{object}	dto.CreateBoardResponse	"Доска успешно создана"
// @Failure		400		{string}	string					"Invalid json payload, Board name is required или ProjectID must be positive"
// @Failure		500		{string}	string					"Cannot create board: <описание ошибки>"
// @Router			/board [post]
func (b *boardHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateBoardRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Board name is required", http.StatusBadRequest)
		return
	}
	if req.ProjectId <= 0 {
		http.Error(w, "ProjectID must be positive", http.StatusBadRequest)
		return
	}

	id, err := b.service.Create(r.Context(), req.Name, req.ProjectId)
	if err != nil {
		http.Error(w, "Cannot create board: "+err.Error(), http.StatusInternalServerError)
		b.logger.Error(r.Context(), "Cannot create board", logger.Field{Key: "error", Value: err.Error()})
		return
	}

	resp := dto.CreateBoardResponse{
		Id:      id,
		Message: "Created board",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		b.logger.Error(r.Context(), "Cannot create board", logger.Field{Key: "error", Value: err.Error()})
	}
}

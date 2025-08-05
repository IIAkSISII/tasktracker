package projectHandler

import (
	"encoding/json"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/transport/http/projectHandler/dto"
	"net/http"
	"time"
)

// @Summary		Создать новый проект
// @Description	Создаем новый проект с указанием названия, даты создания и приватности.
// @Description	Параметры передаются в теле запроса в качестве json-объекта.
// @Description	Если проект не удается создать, возвращаем ошибку.
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			input	body		dto.CreateProjectRequest	true	"Name, isPublic и CreatedAt"
// @Success		201		{object}	dto.CreateProjectResponse	"Проект успешно создан"
// @Failure		400		{string}	string						"Invalid json payload или Project name is required"
// @Failure		500		{string}	string						"Failed to create project: <описание ошибки>"
// @Router			/project [post]
func (p *projectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProjectRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid json payload", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "Project name is required", http.StatusBadRequest)
		return
	}

	if req.CreatedAt.IsZero() {
		req.CreatedAt = time.Now()
	}

	id, err := p.service.Create(r.Context(), req.Name, req.IsPublic, req.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to create project: "+err.Error(), http.StatusInternalServerError)
		p.logger.Error(r.Context(), "Failed to create project", logger.Field{Key: "error", Value: err.Error()})
		return
	}

	resp := dto.CreateProjectResponse{
		Id:      id,
		Message: "Project created",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		p.logger.Error(r.Context(), "Failed to create project", logger.Field{Key: "error", Value: err.Error()})
	}
}

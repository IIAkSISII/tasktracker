package projectHandler

import (
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/service/project"
	"github.com/gorilla/mux"
	"net/http"
)

type projectHandler struct {
	service project.ProjectService
	logger  logger.Logger
}

func NewProjectHandler(service project.ProjectService, logger logger.Logger) *projectHandler {
	return &projectHandler{service: service, logger: logger}
}

func (p *projectHandler) ConfigureRoutes(router *mux.Router) {
	sr := router.PathPrefix("/project").Subrouter()

	sr.HandleFunc("", p.CreateProject).Methods(http.MethodPost)
}

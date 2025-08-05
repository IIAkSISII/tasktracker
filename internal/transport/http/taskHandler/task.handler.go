package taskHandler

import (
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/service/task"
	"github.com/gorilla/mux"
)

type taskHandler struct {
	service task.TaskService
	logger  logger.Logger
}

func NewTaskHandler(service task.TaskService, logger logger.Logger) *taskHandler {
	return &taskHandler{service: service, logger: logger}
}

func (t *taskHandler) ConfigureRoutes(router *mux.Router) {
	sr := router.PathPrefix("/task").Subrouter()

	sr.HandleFunc("", t.CreateTask).Methods("POST")
	sr.HandleFunc("/search", t.SearchTask).Methods("GET")
}

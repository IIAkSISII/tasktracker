package userHandler

import (
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/service/user"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	service user.UserService
	logger  logger.Logger
}

func NewUserHandler(service user.UserService, logger logger.Logger) *UserHandler {
	return &UserHandler{service: service, logger: logger}
}

func (u *UserHandler) ConfigureRoutes(router *mux.Router) {
	sr := router.PathPrefix("/user").Subrouter()

	sr.HandleFunc("", u.CreateUser).Methods(http.MethodPost)
}

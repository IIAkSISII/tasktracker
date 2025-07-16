package userHandler

import (
	"github.com/IIAkSISII/tasktracker/internal/service/user"
	"github.com/gorilla/mux"
	"net/http"
)

type UserHandler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (u *UserHandler) ConfigureRoutes(router *mux.Router) {
	sr := router.PathPrefix("/user").Subrouter()

	sr.HandleFunc("", u.CreateUser).Methods(http.MethodPost)
}

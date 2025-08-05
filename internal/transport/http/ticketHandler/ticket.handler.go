package ticketHandler

import (
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/service/ticket"
	"github.com/gorilla/mux"
)

type ticketHandler struct {
	service ticket.TicketService
	logger  logger.Logger
}

func NewTicketHandler(service ticket.TicketService, logger logger.Logger) *ticketHandler {
	return &ticketHandler{service: service, logger: logger}
}

func (t *ticketHandler) ConfigureRoutes(router *mux.Router) {
	sr := router.PathPrefix("/ticket").Subrouter()

	sr.HandleFunc("", t.CreateTicket).Methods("POST")
	sr.HandleFunc("/move", t.MoveTicket).Methods("PUT")
}

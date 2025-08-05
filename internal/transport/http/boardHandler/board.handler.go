package boardHandler

import (
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/service/board"
	"github.com/gorilla/mux"
)

type boardHandler struct {
	service board.BoardService
	logger  logger.Logger
}

func NewBoardHandler(service board.BoardService, logger logger.Logger) *boardHandler {
	return &boardHandler{service: service, logger: logger}
}

func (b *boardHandler) ConfigureRoutes(router *mux.Router) {
	sr := router.PathPrefix("/board").Subrouter()

	sr.HandleFunc("", b.CreateBoard).Methods("POST")
}

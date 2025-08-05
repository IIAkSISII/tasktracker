package ticket

import (
	"context"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
	"github.com/IIAkSISII/tasktracker/internal/repository/ticket"
	"time"
)

type TicketService interface {
	Create(ctx context.Context, name, description string, userId, LabelId, BoardId int) (int, error)
	Move(ctx context.Context, ticketId, newBoardId int) error
}

type ticketService struct {
	repo   ticket.TicketRepository
	logger logger.Logger
}

func NewTicketService(repo ticket.TicketRepository, logger logger.Logger) TicketService {
	return &ticketService{repo: repo, logger: logger}
}

func (t *ticketService) Create(ctx context.Context, name, description string, userId, LabelId, BoardId int) (int, error) {
	ticket := &models.Ticket{
		Name:        name,
		Description: description,
		UserId:      userId,
		LabelId:     LabelId,
		BoardId:     BoardId,
		CreatedAt:   time.Now(),
	}
	return t.repo.Create(ctx, ticket)
}

func (t *ticketService) Move(ctx context.Context, ticketId, newBoardId int) error {
	return t.repo.Move(ctx, ticketId, newBoardId)
}

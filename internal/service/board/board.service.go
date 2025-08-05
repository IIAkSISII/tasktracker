package board

import (
	"context"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
	"github.com/IIAkSISII/tasktracker/internal/repository/board"
	"time"
)

type BoardService interface {
	Create(ctx context.Context, name string, projectId int) (int, error)
}

type boardService struct {
	repo   board.BoardRepository
	logger logger.Logger
}

func NewBoardService(repo board.BoardRepository, logger logger.Logger) BoardService {
	return &boardService{repo: repo, logger: logger}
}

func (b *boardService) Create(ctx context.Context, name string, projectId int) (int, error) {
	board := &models.Board{
		Name:      name,
		CreatedAt: time.Now(),
		ProjectId: projectId,
	}
	return b.repo.Create(ctx, board)
}

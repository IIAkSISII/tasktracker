package task

import (
	"context"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
	"github.com/IIAkSISII/tasktracker/internal/repository/task"
	"time"
)

type TaskService interface {
	Create(ctx context.Context, name, description string, userId, LabelId, TicketId int) (int, error)
	Search(ctx context.Context, name string, labelId, userId int) ([]*models.Task, error)
}

type taskService struct {
	repo   task.TaskRepository
	logger logger.Logger
}

func NewTicketService(repo task.TaskRepository, logger logger.Logger) TaskService {
	return &taskService{repo: repo, logger: logger}
}

func (t *taskService) Create(ctx context.Context, name, description string, userId, LabelId, TicketId int) (int, error) {
	task := &models.Task{
		Name:        name,
		Description: description,
		UserId:      userId,
		LabelId:     LabelId,
		TicketId:    TicketId,
		CreatedAt:   time.Now(),
	}
	return t.repo.Create(ctx, task)
}

func (t *taskService) Search(ctx context.Context, name string, labelId, userId int) ([]*models.Task, error) {
	return t.repo.Search(ctx, name, labelId, userId)
}

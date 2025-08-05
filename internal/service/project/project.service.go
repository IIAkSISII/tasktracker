package project

import (
	"context"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
	"github.com/IIAkSISII/tasktracker/internal/repository/project"
	"time"
)

type ProjectService interface {
	Create(ctx context.Context, name string, isPublic bool, createdAt time.Time) (int, error)
}

type projectService struct {
	repo   project.ProjectRepository
	logger logger.Logger
}

func NewProjectService(repo project.ProjectRepository, logger logger.Logger) ProjectService {
	return &projectService{repo: repo, logger: logger}
}

func (p *projectService) Create(ctx context.Context, name string, isPublic bool, createdAt time.Time) (int, error) {
	project := &models.Project{
		Name:      name,
		IsPublic:  isPublic,
		CreatedAt: createdAt,
	}
	return p.repo.Create(ctx, project)
}

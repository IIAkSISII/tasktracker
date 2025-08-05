package project

import (
	"context"
	"database/sql"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
)

type ProjectRepository interface {
	Create(ctx context.Context, p *models.Project) (int, error)
}

type projectRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewProjectRepository(db *sql.DB, logger logger.Logger) ProjectRepository {
	return &projectRepository{db: db, logger: logger}
}

func (r *projectRepository) Create(ctx context.Context, p *models.Project) (int, error) {
	query := `INSERT INTO projects (name, is_public, created_at) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, p.Name, p.IsPublic, p.CreatedAt).Scan(&id)
	return id, err
}

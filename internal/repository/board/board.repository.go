package board

import (
	"context"
	"database/sql"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
)

type BoardRepository interface {
	Create(ctx context.Context, p *models.Board) (int, error)
}

type boardRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewBoardRepository(db *sql.DB, logger logger.Logger) BoardRepository {
	return &boardRepository{db: db, logger: logger}
}

func (r *boardRepository) Create(ctx context.Context, p *models.Board) (int, error) {
	query := `INSERT INTO boards(name, created_at, project_id_fk) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, p.Name, p.CreatedAt, p.ProjectId).Scan(&id)
	return id, err
}

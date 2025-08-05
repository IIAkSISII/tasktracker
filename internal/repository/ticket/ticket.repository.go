package ticket

import (
	"context"
	"database/sql"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
)

type TicketRepository interface {
	Create(ctx context.Context, t *models.Ticket) (int, error)
	Move(ctx context.Context, ticketId, newBoardId int) error
}

type ticketRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewTicketRepository(db *sql.DB, logger logger.Logger) TicketRepository {
	return &ticketRepository{db: db, logger: logger}
}

func (r *ticketRepository) Create(ctx context.Context, t *models.Ticket) (int, error) {
	query := `INSERT INTO tickets (name, description, created_at, user_id_fk, label_id_fk, board_id_fk) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, t.Name, t.Description, t.CreatedAt, t.UserId, t.LabelId, t.BoardId).Scan(&id)
	return id, err
}

func (r *ticketRepository) Move(ctx context.Context, ticketId, newBoardId int) error {
	query := `UPDATE tickets SET board_id_fk = $1 WHERE id = $2`

	_, err := r.db.ExecContext(ctx, query, newBoardId, ticketId)
	return err
}

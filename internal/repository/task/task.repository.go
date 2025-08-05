package task

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
)

type TaskRepository interface {
	Create(ctx context.Context, t *models.Task) (int, error)
	Search(ctx context.Context, name string, labelId, userId int) ([]*models.Task, error)
}

type taskRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewTaskRepository(db *sql.DB, logger logger.Logger) TaskRepository {
	return &taskRepository{db: db, logger: logger}
}

func (r *taskRepository) Create(ctx context.Context, t *models.Task) (int, error) {
	query := `INSERT INTO tasks (name, description, created_at, user_id_fk, label_id_fk, ticket_id_fk) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, t.Name, t.Description, t.CreatedAt, t.UserId, t.LabelId, t.TicketId).Scan(&id)
	return id, err
}

func (r *taskRepository) Search(ctx context.Context, name string, labelId, userId int) ([]*models.Task, error) {
	query := `SELECT id, name, description, created_at, user_id_fk, label_id_fk, ticket_id_fk FROM tasks WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	if name != "" {
		query += fmt.Sprintf(" AND name ILIKE $%d", argIndex)
		args = append(args, "%"+name+"%")
		argIndex++
	}
	if labelId > 0 {
		query += fmt.Sprintf(" AND label_id_fk = $%d", argIndex)
		args = append(args, labelId)
		argIndex++
	}
	if userId > 0 {
		query += fmt.Sprintf(" AND user_id_fk = $%d", argIndex)
		args = append(args, userId)
		argIndex++
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		t := new(models.Task)
		err := rows.Scan(&t.Id, &t.Name, &t.Description, &t.CreatedAt, &t.UserId, &t.LabelId, &t.TicketId)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

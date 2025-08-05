package user

import (
	"context"
	"database/sql"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, u *models.User) (int, error)
}

type userRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewUserRepository(db *sql.DB, logger logger.Logger) UserRepository {
	return &userRepository{db: db, logger: logger}
}

func (r *userRepository) Create(ctx context.Context, u *models.User) (int, error) {
	query := `INSERT INTO users (login, email, password_hash) VALUES ($1, $2, $3) RETURNING id`

	var id int
	err := r.db.QueryRowContext(ctx, query, u.Login, u.Email, u.PasswordHash).Scan(&id)
	return id, err
}

package user

import (
	"context"
	"database/sql"
	"github.com/IIAkSISII/tasktracker/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, u *models.User) (int64, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, u *models.User) (int64, error) {
	query := `INSERT INTO users (login, email, password_hash) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	err := r.db.QueryRowContext(ctx, query, u.Login, u.Email, u.PasswordHash).Scan(&id)
	return id, err
}

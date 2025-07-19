package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/IIAkSISII/tasktracker/internal/logger"
	"github.com/IIAkSISII/tasktracker/internal/models"
	"github.com/IIAkSISII/tasktracker/internal/repository/user"
)

type UserService interface {
	Create(ctx context.Context, login, email, password string) (int64, error)
}

type userService struct {
	repo   user.UserRepository
	logger logger.Logger
}

func NewUserService(repo user.UserRepository, logger logger.Logger) UserService {
	return &userService{repo: repo, logger: logger}
}

func (u *userService) Create(ctx context.Context, login, email, password string) (int64, error) {
	hashed := hashPassword(password)
	newUser := &models.User{
		Login:        login,
		Email:        email,
		PasswordHash: hashed,
	}
	return u.repo.Create(ctx, newUser)
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

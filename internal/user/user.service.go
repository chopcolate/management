package user

import (
	"context"

	"github.com/chopcolate/management/internal/common/database"
)

type UserService struct {
	repo *UserRepo
}

func NewUserService(db database.DBTX) *UserService {
	userRepo := NewUserRepo(db)
	return &UserService{userRepo}
}

func (service UserService) CreateUser(ctx context.Context, user *User) string {
	return service.repo.CreateUser(ctx, user)
}

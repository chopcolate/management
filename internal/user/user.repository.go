package user

import (
	"context"
	"log"
	"github.com/chopcolate/management/internal/common/database"
	"github.com/chopcolate/management/internal/common/repository"
)

type UserRepo struct {
	*repository.Repository
}

func NewUserRepo(db database.DBTX) *UserRepo {
	r := repository.NewRepository(db)
	return &UserRepo{r}
}

func (repo *UserRepo) CreateUser(ctx context.Context, user *User) string {
	var newUser User
	sql := `INSERT INTO ACCOUNT (name, username, password, phone, email, role) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, name, username, phone, email, role`
	repo.DB.QueryRow(ctx, sql, user.Name, user.Username, user.Password, user.Phone, user.Email, user.Role).Scan(&newUser.ID, &newUser.Name, &newUser.Username, &newUser.Phone, &newUser.Email, &newUser.Role)
	log.Fatal(newUser)
	return "A new user has been created"
}

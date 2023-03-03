package repository

import (
	"github.com/chopcolate/management/internal/common/database"
)

type Repository struct {
	DB database.DBTX
}

func NewRepository(db database.DBTX) *Repository {
	return &Repository{DB: db}
}

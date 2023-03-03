package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	DB *pgxpool.Pool
}

func NewDatabase(ctx context.Context, url string) (*Database, error) {
	db, err := pgxpool.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	return &Database{DB: db}, nil
}

type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

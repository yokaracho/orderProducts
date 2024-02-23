package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

func ConnectionPool(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	pool, err := pgxpool.Connect(ctx, connString)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}

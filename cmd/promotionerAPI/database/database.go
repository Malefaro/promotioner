package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabase(addr, name, username, password string) (*Database, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=verify-full", username, password, addr, name))
	if err != nil {
		return nil, err
	}
	config.MaxConns = 6
	config.MinConns = 2
	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Database{Pool: conn}, nil
}

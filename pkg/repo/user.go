package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type User struct {
	Id int
	Name string
}

type UserStorage struct {
	pool *pgxpool.Pool
}

func NewUserStorage(pool *pgxpool.Pool) *UserStorage {
	return &UserStorage{
		pool: pool,
	}
}

func (us *UserStorage) CreateUser(ctx context.Context, uName string) (int, error) {
	var id int
  	const query = `INSERT INTO users (name) VALUES ($1) RETURNING id;`

	err := us.pool.QueryRow(ctx, query, uName).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}
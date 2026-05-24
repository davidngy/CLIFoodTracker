package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
}

func CreateUser(ctx context.Context, db *pgxpool.Pool, name string) (User, error) {
	var user User

	query := `
			INSERT INTO users (name)
			VALUES ($1)
			RETURNING id, name, created_at
	`

	err := db.QueryRow(ctx, query, name).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
	)
	if err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

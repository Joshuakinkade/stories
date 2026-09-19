package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type InsertUserParams struct {
	ID          uuid.UUID
	Email       string
	DisplayName string
	Slug        string
	Password    string
	CreatedAt   time.Time
	UpdatedAT   time.Time
}

const insertStmt string = "INSERT INTO users(id, email, display_name, slug, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7);"

func InsertUser(ctx context.Context, user InsertUserParams) error {
	conn, err := GetConnection(ctx)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, insertStmt, user.ID, user.Email, user.DisplayName, user.Slug, user.Password)
	return err
}

func FindUserById(ctx context.Context, userId uuid.UUID) error {
	return nil
}

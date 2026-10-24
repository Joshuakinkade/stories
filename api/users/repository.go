package users

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/joshuakinkade/web-journal/db"
)

type Repository struct{}

func NewRepository() Repository {
	return Repository{}
}

const insertStmt string = "INSERT INTO users(id, email, display_name, slug, password, roles, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8);"

func (r Repository) Create(ctx context.Context, u *User) error {
	u.updatedAt = time.Now()

	conn, err := db.GetConnection(ctx)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, insertStmt, u.ID(), u.Email(), u.DisplayName(), u.Slug(), u.PasswordHash(), getStringRoles(u.roles), u.CreatedAt(), u.UpdatedAt())
	return err
}

func getStringRoles(roles []UserRole) []string {
	var roleStrings []string
	for _, r := range roles {
		roleStrings = append(roleStrings, r.String())
	}
	return roleStrings
}

const updateStmt string = "UPDATE users SET email = $2, display_name = $3, slug = $4, password = $5, updated_at = $6 WHERE id = $1"

func (r Repository) Update(ctx context.Context, u *User) error {
	u.updatedAt = time.Now()

	conn, err := db.GetConnection(ctx)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, updateStmt, u.ID(), u.Email(), u.DisplayName(), u.Slug(), u.PasswordHash(), u.UpdatedAt())
	return err
}

const findByIDStmt string = "SELECT id, email, display_name, slug, password, created_at, updated_at FROM users WHERE id = $1"

func (r Repository) FindByID(ctx context.Context, targetID uuid.UUID) (User, error) {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return User{}, err
	}

	return scanUser(conn.QueryRow(ctx, findByIDStmt, targetID))
}

const findByEmailStmt string = "SELECT id, email, display_name, slug, password, created_at, updated_at FROM users WHERE email = $1"

func (r Repository) FindByEmail(ctx context.Context, targetEmail string) (User, error) {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return User{}, err
	}

	return scanUser(conn.QueryRow(ctx, findByEmailStmt, targetEmail))
}

// scanUser reads a row selected as: id, email, display_name, slug, password, created_at, updated_at.
func scanUser(row pgx.Row) (User, error) {
	var u User
	err := row.Scan(&u.id, &u.email, &u.displayName, &u.slug, &u.password, &u.createdAt, &u.updatedAt)
	if err != nil {
		return User{}, err
	}

	return u, nil
}

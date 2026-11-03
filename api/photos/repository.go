package photos

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

const photoColumns string = "id, user_id, alt_text, caption, height, width, created_at, updated_at"

const insertStmt string = "INSERT INTO photos(id, user_id, alt_text, caption, height, width, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"

func (r Repository) Create(ctx context.Context, p *Photo) error {
	p.updatedAt = time.Now()

	conn, err := db.GetConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, insertStmt, p.id, p.userID, p.altText, p.caption, p.height, p.width, p.createdAt, p.updatedAt)
	return err
}

const updateStmt string = "UPDATE photos SET alt_text = $2, caption = $3, height = $4, width = $5, updated_at = $6 WHERE id = $1"

func (r Repository) Update(ctx context.Context, p *Photo) error {
	p.updatedAt = time.Now()

	conn, err := db.GetConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, updateStmt, p.id, p.altText, p.caption, p.height, p.width, p.updatedAt)
	return err
}

const findByIDStmt string = "SELECT " + photoColumns + " FROM photos WHERE id = $1"

func (r Repository) FindByID(ctx context.Context, photoID uuid.UUID) (Photo, error) {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return Photo{}, err
	}
	defer conn.Release()

	return scanPhoto(conn.QueryRow(ctx, findByIDStmt, photoID))
}

const listByUserStmt string = "SELECT " + photoColumns + " FROM photos WHERE user_id = $1 ORDER BY created_at DESC"

func (r Repository) ListByUser(ctx context.Context, userID uuid.UUID) ([]Photo, error) {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, listByUserStmt, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []Photo
	for rows.Next() {
		p, err := scanPhoto(rows)
		if err != nil {
			return nil, err
		}
		photos = append(photos, p)
	}
	return photos, rows.Err()
}

const deleteStmt string = "DELETE FROM photos WHERE id = $1"

func (r Repository) Delete(ctx context.Context, photoID uuid.UUID) error {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, deleteStmt, photoID)
	return err
}

// scanPhoto reads a row selected as: id, user_id, alt_text, caption, height, width, created_at, updated_at.
func scanPhoto(row pgx.Row) (Photo, error) {
	var p Photo
	err := row.Scan(&p.id, &p.userID, &p.altText, &p.caption, &p.height, &p.width, &p.createdAt, &p.updatedAt)
	if err != nil {
		return Photo{}, err
	}
	return p, nil
}

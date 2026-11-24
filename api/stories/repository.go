package stories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IRepository interface {
	GetStoryByID(id uuid.UUID) (Story, error)
	SaveStory(story Story) error
}

type Repository struct {
	dbPool *pgxpool.Pool
}

func (r Repository) GetStoryByID(id uuid.UUID) (Story, error) {
	return Story{}, nil
}

func (r Repository) SaveStory(story Story) error {
	conn, err := r.dbPool.Acquire(context.TODO())
	if err != nil {
		return err
	}
	if story.isNew {
		_, err := conn.Exec(context.TODO(), "", nil)
		return err
	} else {
		_, err := conn.Exec(context.TODO(), "", nil)
		return err
	}
}

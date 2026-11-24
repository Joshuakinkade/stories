package stories

import (
	"time"

	"github.com/google/uuid"
)

type Version struct {
	id          uuid.UUID
	storyID     uuid.UUID
	slug        string
	publishedAt *time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func NewVersion(slug string, storyID uuid.UUID) Version {
	id, _ := uuid.NewV7()
	return Version{
		id:          id,
		storyID:     storyID,
		slug:        slug,
		publishedAt: nil,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
	}
}

func (v Version) ID() uuid.UUID {
	return v.id
}

func (v Version) StoryID() uuid.UUID {
	return v.storyID
}

func (v Version) Slug() string {
	return v.slug
}

func (v Version) PublishedAt() *time.Time {
	return v.publishedAt
}

func (v Version) IsPublished() bool {
	return v.publishedAt != nil
}

func (v *Version) Publish() {
	t := time.Now()
	v.publishedAt = &t
}

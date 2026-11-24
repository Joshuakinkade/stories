package stories

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Story struct {
	id        uuid.UUID
	title     string
	slug      string
	elements  []Element
	createdAt time.Time
	updatedAt time.Time

	isNew bool
}

func NewStory(id uuid.UUID, title string) Story {
	return Story{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		isNew:     true,
	}
}

func FindByID(id uuid.UUID) (Story, error) {
	return Story{}, nil
}

func Restore(id uuid.UUID, title, slug string, publishedAt *time.Time, createdAt, updatedAt time.Time) Story {
	return Story{
		id:        id,
		title:     title,
		slug:      slug,
		createdAt: createdAt,
		updatedAt: updatedAt,

		isNew: false,
	}
}

func (s Story) ID() uuid.UUID {
	return s.id
}

func (s Story) Title() string {
	return s.title
}

func (s *Story) SetTitle(title string) {
	s.title = title
	s.slug = title
}

func (s *Story) RefershSlug() {
	slug := strings.ToLower(s.title)
	s.slug = strings.ReplaceAll(slug, " ", "-")
}

func (s Story) Slug() string {
	return s.slug
}

func (s *Story) AddElement(el Element) {
	s.elements = append(s.elements, el)
}

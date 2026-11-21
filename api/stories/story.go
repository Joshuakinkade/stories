package stories

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Story struct {
	id          uuid.UUID
	title       string
	slug        string
	elements    []Element
	publishedAt *time.Time
	createdAt   time.Time
	updatedAt   time.Time
}

func NewStory(id uuid.UUID, title string) Story {
	return Story{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func FindByID(id uuid.UUID) (Story, error) {
	return Story{}, nil
}

func Restore(id uuid.UUID, title, slug string, publishedAt *time.Time, createdAt, updatedAt time.Time) Story {
	return Story{
		id:          id,
		title:       title,
		slug:        slug,
		publishedAt: publishedAt,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
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

func (s *Story) IsPublished() bool {
	return s.publishedAt != nil
}

func (s Story) PublishedAt() time.Time {
	return *s.publishedAt
}

func (s *Story) CreateVersion() error {
	return nil
}

func (s *Story) Publish() error {
	t := time.Now()
	s.publishedAt = &t

	// create a new version
	//   copy config and elements to new version
	// pin the published version
	// do this in a transaction
	return nil
}

func (s *Story) Save() error {
	s.updatedAt = time.Now()
	return nil
}

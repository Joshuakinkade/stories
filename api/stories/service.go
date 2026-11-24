// Package stories manages a user's stories.
package stories

import (
	"github.com/google/uuid"
)

type StoryUpdates struct {
	Title *string
}

type Service struct {
	repo IRepository
}

func NewService() Service {
	return Service{}
}

// CreateStory creates a new story for a user.
func (s Service) CreateStory() error {
	// Create the story struct
	// Create the first version struct
	return nil
}

// UpdateStory patches an existing story's basic information.
func (s Service) UpdateStory(slug string, updates StoryUpdates) error {
	story, err := s.resolveWorkingStory(slug)
	if err != nil {
		return nil
	}

	// do some updates here.
	if updates.Title != nil {
		story.SetTitle(*updates.Title)
	}

	return s.repo.SaveStory(story)
}

// DeleteStory permanently deletes a story.
func (s Service) DeleteStory() error {
	return nil
}

// AddElement adds an element to a story at the given index.
func (s Service) AddElement() error {
	return nil
}

// MoveElement moves an element to a new location.
func (s Service) MoveElement() error {
	return nil
}

// RemoveElement removes an element from a story.
func (s Service) RemoveElement() error {
	return nil
}

// PublishStory takes a snapshot of story and makes it available in their
// gallery.
func (s Service) PublishStory(storyID uuid.UUID) error {
	return nil
}

// GetStoryBySlug fetches the publised version of story by it's slug.
func (s Service) GetStoryBySlug(slug string) (Story, error) {
	return Story{}, nil
}

func (s Service) resolveWorkingStory(slug string) (Story, error) {
	return Story{}, nil
}

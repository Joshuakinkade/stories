// Package stories manages a user's stories.
package stories

import (
	"github.com/google/uuid"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

// CreateStory creates a new story for a user.
func CreateStory() error {
	return nil
}

// UpdateStory patches an existing story's basic information.
func UpdateStory() error {
	return nil
}

// DeleteStory permanently deletes a story.
func DeleteStory() error {
	return nil
}

// AddElement adds an element to a story at the given index.
func AddElement() error {
	return nil
}

// MoveElement moves an element to a new location.
func MoveElement() error {
	return nil
}

// RemoveElement removes an element from a story.
func RemoveElement() error {
	return nil
}

// PublishStory takes a snapshot of story and makes it available in their
// gallery.
func PublishStory(storyID uuid.UUID) error {
	s, err := FindByID(storyID)
	if err != nil {
		return err
	}

	err = s.Publish()
	if err != nil {
		return err
	}
	return nil
}

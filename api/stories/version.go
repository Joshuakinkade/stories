package stories

import (
	"time"

	"github.com/google/uuid"
)

type Version struct {
	id        uuid.UUID
	elements  []Element
	createdAt time.Time
	updatedAt time.Time
}

func NewVersion(elements []Element) Version {
	id, _ := uuid.NewV7()
	return Version{
		id:        id,
		elements:  elements,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

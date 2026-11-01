package photos

import (
	"json"
	"time"

	"github.com/google/uuid"
)

type Album struct {
	id uuid.UUID
	name string
	createdAt time.Time
	updatedAt time.Time
}

func NewAlbum(name string) Album {
	id, _ := uuid.NewV7()
	return Album{
		id: id,
		name: name,
		createdAt: time.Now(),
	}
}

func (a Album) ID() uuid.UUID {
	return a.id
}

func (a Album) Name() string {
	return a.name
}

func (a *Album) SetName(newName string) {
	a.name = newName
	a.updatedAt = time.Now()
}

func (a Album) CreatedAt() time.Time {
	return a.createdAt
}

func (a Album) UpdatedAt() time.Time {
	return a.updatedAt
}

func (a Album) MarshalJSON() ([]byte, error) {
	dataMap := map[string]any{
		"id": a.ID(),
		"name": a.Name(),
		"createdAt": a.CreatedAt(),
		"updatedAt": a.UpdatedAt()k
	}

	return json.Marshal(dataMap)
}

package photos

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	id        uuid.UUID
	userID    uuid.UUID
	altText   string
	caption   string
	height    int
	width     int
	createdAt time.Time
	updatedAt time.Time

	isNew bool
}

func NewPhoto(userID uuid.UUID) Photo {
	id, _ := uuid.NewV7()
	createdAt := time.Now()
	p := Photo{
		id:        id,
		userID:    userID,
		createdAt: createdAt,
	}

	return p
}

func (p Photo) ID() uuid.UUID {
	return p.id
}

func (p Photo) UserID() uuid.UUID {
	return p.userID
}

func (p Photo) CreatedAt() time.Time {
	return p.createdAt
}

func (p Photo) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p Photo) AltText() string {
	return p.altText
}

func (p *Photo) SetAltText(newAltText string) {
	p.altText = newAltText
}

func (p Photo) Caption() string {
	return p.caption
}

func (p *Photo) SetCaption(newCaption string) {
	p.caption = newCaption
}

func (p Photo) Dimensions() (int, int) {
	return p.height, p.width
}

func (p *Photo) SetDimensions(height, width int) {
	p.height = height
	p.width = width
}

package photos

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	id        uuid.UUID
	altText   string
	caption   string
	height    int
	width     int
	createdAt time.Time
	updatedAt time.Time

	isNew bool
}

func NewPhoto() Photo {
	id, _ := uuid.NewV7()
	createdAt := time.Now()
	p := Photo{
		id:        id,
		createdAt: createdAt,
	}

	return p
}

func (p Photo) ID() uuid.UUID {
	return uuid.UUID{}
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

func (p *Photo) Save() error {
	p.updatedAt = time.Now()

	if p.isNew {
		// insert
	} else {
		// update
	}

	return nil
}

func FindPhoto() (Photo, error) {
	return Photo{}, nil
}

/************************************** QUERY METHODS BELOW **************************************/

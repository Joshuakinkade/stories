package stories

import "github.com/google/uuid"

type ElementConfig map[string]any

type Element struct {
	id          uuid.UUID
	elementType string
	config      ElementConfig
}

func NewElement(t string) Element {
	id, _ := uuid.NewV7()
	return Element{
		id:          id,
		elementType: t,
	}
}

func (e Element) ID() uuid.UUID {
	return e.id
}

func (e Element) Config() ElementConfig {
	return e.config
}

func (e *Element) SetConfigValue(key string, value any) {
	e.config[key] = value
}

func (e Element) MarshalJSON() ([]byte, error) {
	return nil, nil
}

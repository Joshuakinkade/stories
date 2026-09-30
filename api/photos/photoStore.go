package photos

import "github.com/google/uuid"

type Store struct {
}

func NewStore() Store {
	return Store{}
}

func (s Store) Save(tempFileName string, photoId uuid.UUID, albumId uuid.UUID) error {
	return nil
}

func (s Store) Get()

package photos

type Store struct {

}

func NewStore() Store {

}

func (s Store) Save(tempFileName string, photoId uuid.UUID, albumId uuid.UUID) error {
  return nil
}

func (s Store) Get()

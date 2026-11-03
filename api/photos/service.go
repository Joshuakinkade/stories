package photos

import (
 	"context"

	"github.com/google/uuid"
)

type PhotoService struct {
	store Store
	repository Repository
}

func NewPhotoService(store Store) PhotoService {
	return {
		store: store
	}
}

func (ps PhotoService) UploadPhoto(ctx context.Context) error {
	p := NewPhoto()

	// check photo dimensions
	// load into Image

	// reorient based on exif data or something else
	// resize if too big, maybe 2mb
	// save to photo store with original format
	err := ps.store.Save(ctx, []byte{}, p.ID())
	if err != nil {
		return errors.New("Could not save photo to store")
	}

	// extract some metadata: height, width, date taken, location, etc.
	// save metadata to photo struct
	// save photo
	err := ps.repository.Create(ctx, p)
	if err != nil {
		return errors.New("Could not save photo to db")
	}

	return nil
}

func (ps PhotoService) ListPhotos(ctx context.Context, userID uuid.UUID) ([]Photo, error) {
	return nil, nil
}

func (ps PhotoService) GetPhoto(ctx context.Context, photoID uuid.UUID) (Photo, error) {
	return Photo{}, nil
}

func (ps PhotoService) ListAlbums(ctx context.Context, userID uuid.UUID) ([]Album, error) {
	return nil, nil
}

func (ps PhotoService) CreateAlbum(ctx context.Context, userID uuid.UUID, name string) error {
	return nil
}

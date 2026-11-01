package photos

import (
 	"context"

	"github.com/google/uuid"
)

type PhotoService struct {
}

func (ps PhotoService) UploadPhoto(ctx context.Context) error {
	// p := NewPhoto()

	// check photo dimensions
	// load into Image

	// reorient based on exif data or something else
	// resize if too big, maybe 2mb
	// save to photo store with original format

	// extract some metadata: height, width, date taken, location, etc.
	// save metadata to photo struct
	// save photo

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

package photos

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"

	"github.com/google/uuid"
	"github.com/disintegration/imaging"
)

type Service struct {
	store      Store
	repository Repository
}

func NewService(store Store) Service {
	return Service{
		store: store,
	}
}

func (ps Service) UploadPhoto(ctx context.Context, userID uuid.UUID, file io.Reader, fileName string) (Photo, error) {
	
	p := NewPhoto(userID)

	img, err := imaging.Decode(file, imaging.AutoOrientation(true))
	if err != nil {
		log.Println(err)
		return Photo{}, errors.New("could not read image")
	}

	// get some metadata, save to photo struct 
 
	// save to object storage
  var b bytes.Buffer
	err = imaging.Encode(&b, img, imaging.JPEG)
	if err != nil {
		return Photo{}, err
	}
	err = ps.store.Save(ctx, b.Bytes(), p.ID())
	if err != nil {
		return Photo{}, errors.New("Could not save photo to store")
	}

	err = ps.repository.Create(ctx, &p)
	if err != nil {
		return Photo{}, errors.New("Could not save photo to db")
	}

	return p, nil
}

func (ps Service) ListPhotos(ctx context.Context, userID uuid.UUID) ([]Photo, error) {
	return nil, nil
}

func (ps Service) GetPhoto(ctx context.Context, photoID uuid.UUID) ([]byte, error) {
	p, err := ps.store.Get(ctx, photoID)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (ps Service) ListAlbums(ctx context.Context, userID uuid.UUID) ([]Album, error) {
	return nil, nil
}

func (ps Service) CreateAlbum(ctx context.Context, userID uuid.UUID, name string) error {
	return nil
}

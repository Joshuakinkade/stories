// Package photos handles saving, processing, and retrieving photos for the application
package photos

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
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
		return Photo{}, errors.New("could not save photo to store")
	}

	err = ps.repository.Create(ctx, &p)
	if err != nil {
		return Photo{}, errors.New("could not save photo to db")
	}

	return p, nil
}

func (ps Service) ListPhotos(ctx context.Context, userID uuid.UUID) ([]Photo, error) {
	return nil, nil
}

// OutputParams defines options for the returned photo
type OutputParams struct {
	Format  *string
	Height  *int
	Width   *int
	Quality *int
}

// GetPhoto returns a copy of the photo with the given ids that has been to fit the given
// parameters. If the variant has already been created, an existing copy wetll be returned.
func (ps Service) GetPhoto(ctx context.Context, photoID uuid.UUID, params OutputParams) ([]byte, error) {
	variantID := ps.buildPhotoVariantID(photoID, params)
	log.Println(variantID)

	// try to get variant from photo store
	// if found, return it
	// else if variantID != photoID
	//  get master photo, return if not found
	//  process it
	//  save variant
	//  return variant
	// else return not found

	p, err := ps.store.Get(ctx, photoID)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (ps Service) buildPhotoVariantID(photoID uuid.UUID, params OutputParams) string {
	var parts []string
	if params.Format != nil {
		parts = append(parts, fmt.Sprintf(";format:%v", params.Format))
	}
	if params.Height != nil {
		parts = append(parts, fmt.Sprintf(";height:%v", params.Height))
	}
	if params.Width != nil {
		parts = append(parts, fmt.Sprintf(";width:%v", params.Width))
	}
	if params.Quality != nil {
		parts = append(parts, fmt.Sprintf(";quality:%v", params.Quality))
	}

	id := photoID.String()
	if len(parts) > 0 {
		suffix := strings.Join(parts, ";")
		encoded := make([]byte, hex.EncodedLen(len(suffix)))
		hex.Encode([]byte(id), encoded)
		return id + ":" + string(encoded)
	}

	return id
}

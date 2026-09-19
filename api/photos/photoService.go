package photos

type PhotoService struct {

}

func (ps PhotoService) UploadPhoto() error {
	p := NewPhoto()

	// open temp photo file
	
	// reorient based on exif data or something else
	// resize if too big, maybe 2mb
	// save to photo store
	
	// extract some metadata: height, width, date taken, location, etc.
	// save metadata to photo struct
  // save photo
	
	return nil
}

func (ps PhotoService) GetPhotos(userID uuid.UUID) ([]Photo, error) {
	return nil, nil
}

func (ps PhotoService) GetPhoto(photoID uuid.UUID) (Photo, error) {
	return Photo{}, nil
}

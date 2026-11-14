package photos

import (
	"errors"
)

var ErrInvalidFormat = errors.New("invalid image format")

type Format int

func (f Format) MimeType() (string, error) {
	switch f {
	case JPEG:
		return "image/jpeg", nil
	case WEBP:
		return "image/webp", nil
	case HEIC:
		return "image/heic", nil
	case PNG:
		return "image/png", nil
	}
	return "", ErrInvalidFormat
}

const (
	JPEG Format = iota
	WEBP
	HEIC
	PNG
)

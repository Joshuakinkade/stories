package web

import (
	"github.com/gofiber/fiber/v3"
	"github.com/joshuakinkade/web-journal/photos"
)

type PhotoHandlers struct {
	photos photos.Service
}

func NewPhotoHandlers(p photos.Service) PhotoHandlers {
	return PhotoHandlers{
		photos: p,
	}
}

func (p PhotoHandlers) Upload(ctx fiber.Context) error {

	return nil
}

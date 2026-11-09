package web

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
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

func (p PhotoHandlers) Upload(ctx fiber.Ctx) error {
  userID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		return errors.New("bad userID")
	}

	fileHeader, err := ctx.FormFile("photo")
	if err != nil {
		return err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return err
	}

	photo, err := p.photos.UploadPhoto(ctx, userID, file, fileHeader.Filename)
	if err != nil {
		ctx.JSON(map[string]string {
			"error": err.Error(),
		})
		return nil
	} else {
		ctx.JSON(photo)
	}

	return nil
}

func (p PhotoHandlers) GetPhoto(ctx fiber.Ctx) error {
	photoIDStr := ctx.Params("photo_id")
	photoID, err := uuid.Parse(photoIDStr)
	if err != nil {
		return ctx.JSON(map[string]string{
			"error": "photoID invalid",
		})
	}

	photo, err := p.photos.GetPhoto(ctx, photoID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest)
		ctx.JSON(map[string]string {
			"error": "could not get photo",
		})
		return nil
	}

	ctx.Type("jpeg")
	return ctx.Send(photo)
}

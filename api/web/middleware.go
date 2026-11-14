package web

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/joshuakinkade/web-journal/auth"
)

type Middleware struct {
	auth auth.Manager
}

func NewMiddleware(a auth.Manager) Middleware {
	return Middleware{
		auth: a,
	}
}

func (m Middleware) AuthMiddleware(c fiber.Ctx) error {
	const prefix = "Bearer "
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		c.Status(fiber.StatusUnauthorized)
		_, err := c.Write([]byte("Authorization header is required"))
		return err
	}
	if !strings.HasPrefix(authHeader, prefix) {
		c.Status(fiber.StatusUnauthorized)
		_, err := c.WriteString("Authorization header is not valid")
		return err
	}
	token := strings.TrimPrefix(authHeader, prefix)

	userIDstr, authErr := m.auth.CheckJWT(token)
	if authErr != nil {
		c.Status(fiber.StatusUnauthorized)
		_, err := c.WriteString(authErr.Error())
		return err
	}
	userID, err := uuid.Parse(userIDstr)
	if err != nil {
		return errors.New("could no parse error")
	}

	c.Locals("userID", userID)
	return c.Next()
}

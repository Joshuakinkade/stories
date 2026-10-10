package web

import (
	"strings"

	"github.com/gofiber/fiber/v3"
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
		c.Status(401)
		c.Write([]byte("Authorization header is required"))
		return nil
	}
	if !strings.HasPrefix(authHeader, prefix) {
		c.Status(401)
		c.WriteString("Authorization header is not valid")
		return nil
	}
	token := strings.TrimPrefix(authHeader, prefix)

	_, authErr := m.auth.CheckJWT(token)
	if authErr != nil {
		c.Status(401)
		c.WriteString(authErr.Error())
		return nil
	}
	return c.Next()
}

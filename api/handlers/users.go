package handlers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/joshuakinkade/web-journal/dtos"
	"github.com/joshuakinkade/web-journal/users"
)

type UserHandlers struct {
	users users.Service
}

func (h UserHandlers) Register(ctx fiber.Ctx) error {
	newUser := dtos.NewUserInput{
		Email:           ctx.FormValue("email"),
		Password:        ctx.FormValue("password"),
		ConfirmPassword: ctx.FormValue("confirm_password"),
	}

	u, err := h.users.CreateUser(newUser, false)
	if err != nil { // Check for error type yere
		return ctx.JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(u)
}

type LoginResponse struct {
	User users.User `json:"user"`
	Token string    `json:"token"`
	Error string    `json:"error"`
}

func (h UserHandlers) Login(ctx fiber.Ctx) error {

	email := ctx.FormValue("email")
	if email == "" {
		ctx.Status(http.StatusBadRequest)
		ctx.WriteString("Email is required")
		return nil
	}
	password := ctx.FormValue("password")
	if password == "" {
		ctx.Status(http.StatusBadRequest)
		ctx.WriteString("Password is required")
		return nil
	}

	u, authToken, _, err := h.users.LoginUser(email, password)
	if err != nil {
		log.Println(err)
		ctx.Status(http.StatusBadRequest)
		ctx.JSON(LoginResponse{
			Error: "Could not log in user",
		})
		return nil
	}

	ctx.Status(200)
	ctx.JSON(LoginResponse{
		User: u,
		Token: authToken,
	})
	return nil
}

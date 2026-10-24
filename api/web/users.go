package web

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/joshuakinkade/web-journal/dtos"
	"github.com/joshuakinkade/web-journal/users"
)

type UserHandlers struct {
	users users.Service
}

func NewUserHandlers(u users.Service) UserHandlers {
	return UserHandlers{
		users: u,
	}
}

type LoginResponse struct {
	User  users.User `json:"user"`
	Token string     `json:"token"`
	Error string     `json:"error"`
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
		User:  u,
		Token: authToken,
	})
	return nil
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

func (h UserHandlers) GetUser(ctx fiber.Ctx) error {
	userId, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		return errors.New("Invalid User ID")
	}

	user, err := h.users.GetByUserID(userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	} else if errors.Is(err, sql.ErrNoRows) {
		return ctx.JSON(map[string]string{
			"error": "User not found",
		})
	}

	return ctx.JSON(user)
}

func (h UserHandlers) UpdateUser(ctx fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		return errors.New("Invalid User ID")
	}

	update := dtos.UserUpdate{}

	displayName := ctx.FormValue("display_name")
	update.DisplayName = &displayName

	err := h.users.UpdateUser(userID, update)
	if err != nil {
		return err
	}

	ctx.Status(fiber.StatusOK)
	ctx.WriteString("Updated")
	return nil
}

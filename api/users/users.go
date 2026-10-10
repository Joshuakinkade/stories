package users

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/joshuakinkade/web-journal/auth"
	"github.com/joshuakinkade/web-journal/dtos"
	"github.com/joshuakinkade/web-journal/validation"
)

type Service struct {
	auth auth.Manager
}

func NewService(a auth.Manager) Service {
	return Service{
		auth: a,
	}
}

func (s *Service) CreateUser(newUser dtos.NewUserInput, isAdmin bool) (User, error) {
	err := validateUser(newUser)
	if err != nil {
		return User{}, err
	}

	// Create User
	u := NewUser(newUser.Email, newUser.DisplayName, newUser.Password)

	u.AddRole(UserAuthorRole)
	if isAdmin {
		u.AddRole(UserAdminRole)
	}

	err = u.Save()
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *Service) CreateUserWithCode(newUser dtos.NewUserInput, code string) (User, error) {
	err := validateUser(newUser)
	if err != nil {
		return User{}, err
	}

	// Check registration code
	rc, err := GetByCode(code)
	if err != nil {
		return User{}, nil
	}
	if rc.IsExpired() {
		return User{}, errors.New("registration code is expired")
	}

	// Create User
	u := NewUser(newUser.Email, newUser.DisplayName, newUser.Password)
	u.AddRole(UserAuthorRole)
	err = u.Save()
	if err != nil {
		return User{}, err
	}

	// Create JWT here

	return u, nil
}

func validateUser(newUser dtos.NewUserInput) error {
	errors := map[string]string{}

	if err := validation.ValidateEmail(newUser.Email); err != nil {
		errors["email"] = err.Error()
	}
	if newUser.Password != newUser.ConfirmPassword {
		errors["password"] = "passwords don't match"
	}
	if len(errors) != 0 {
		return validation.Error{
			Errors: errors,
		}
	}
	return nil
}

func (s Service) LoginUser(email, password string) (User, string, string, error) {
	user, err := FindByEmail(context.Background(), email)
	if user.ID() == uuid.Nil {
		return User{}, "", "", err
	}
	err = user.CheckPassword(password)
	if err != nil {
		return User{}, "", "", errors.New("password invalid")
	}

	authToken, err := s.auth.CreateAccessToken(user.ID())
	if err != nil {
		return User{}, "", "", err
	}

	refreshToken, err := s.auth.CreateRefreshToken(user.ID())
	if err != nil {
		return User{}, "", "", err
	}

	return user, authToken, refreshToken, nil
}

func (s Service) ListUsers() ([]User, error) {
	return nil, nil
}

// DO THIS LATER, need to think about how to pass typed updates
// type UpdateField string

// const DisplayName UpdateField = "displayName"
// const Password UpdateField = "password"

// func (s UserService) UpdateUser(id uuid.UUID, updates map[UpdateField]any) error {
// 	u, err := FindByID(id)
// 	if err != nil { // Also need to check for not found
// 		return err
// 	}

// 	if value, ok := updates[DisplayName]; ok {

// 	}

// 	return nil
// }

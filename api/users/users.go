package users

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/joshuakinkade/web-journal/auth"
	"github.com/joshuakinkade/web-journal/dtos"
	"github.com/joshuakinkade/web-journal/validation"
)

type Service struct {
	auth auth.Manager
	repo Repository
}

func NewService(a auth.Manager, r Repository) Service {
	return Service{
		auth: a,
		repo: r,
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

	err = s.repo.Create(context.Background(), &u)
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
	err = s.repo.Create(context.Background(), &u)
	if err != nil {
		return User{}, err
	}

	// Create JWT here

	return u, nil
}

func validateUser(newUser dtos.NewUserInput) error {
	errors := map[string]string{}

	if _, err := validation.ValidateEmail(newUser.Email); err != nil {
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
	user, err := s.repo.FindByEmail(context.Background(), email)
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

func (s Service) GetByUserID(userID uuid.UUID) (User, error) {
	user, err := s.repo.FindByID(context.Background(), userID)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func (s Service) UpdateUser(userId uuid.UUID, update dtos.UserUpdate) error {
	u, err := s.repo.FindByID(context.Background(), userId)
	if err != nil {
		return err
	}

	if update.Email != nil {
		u.SetEmail(*update.Email)
	}

	if update.DisplayName != nil {
		u.SetDisplayName(*update.DisplayName)
	}


	if update.Password != nil && update.ConfirmPassword != nil {
		log.Println("%v %v", *update.Password, *update.ConfirmPassword)
		if *update.Password == *update.ConfirmPassword {
			u.SetPassword(*update.Password)
		}
	}

	return s.repo.Update(context.Background(), &u)
	return nil
}

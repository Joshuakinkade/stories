package users

import (
	"time"

	"github.com/google/uuid"
)

type RegistrationCode struct {
	id        uuid.UUID
	code      string
	expiresAt time.Time
	createdAt time.Time
}

func NewRegistrationCode(code string) RegistrationCode {
	id, _ := uuid.NewV7()
	return RegistrationCode{
		id:        id,
		code:      code,
		expiresAt: time.Now().Add(time.Hour * 24 * 7), // Expire after a week
	}
}

func GetByCode(code string) (RegistrationCode, error) {
	return RegistrationCode{}, nil
}

func (rc RegistrationCode) ID() uuid.UUID {
	return rc.id
}

func (rc RegistrationCode) Code() string {
	return rc.code
}

func (rc RegistrationCode) ExpiresAt() time.Time {
	return rc.expiresAt
}

func (rc RegistrationCode) IsExpired() bool {
	return rc.expiresAt.After(time.Now())
}

func (rc RegistrationCode) CreatedAt() time.Time {
	return rc.createdAt
}

func (rc RegistrationCode) Save() error {
	return nil
}

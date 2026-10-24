package users

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserRole int

func (r UserRole) String() string {
	switch r {
	case UserAuthorRole:
		return "UserAuthorRole"
	case UserAdminRole:
		return "UserAdminRole"
	default:
		return "<INVALID>"
	}
}

const (
	UserAuthorRole UserRole = iota
	UserAdminRole
)

type User struct {
	id          uuid.UUID
	email       string
	displayName string
	slug        string
	password    string
	roles       []UserRole
	createdAt   time.Time
	updatedAt   time.Time

	rolesAdded   []uuid.UUID
	rolesRemoved []uuid.UUID
}

func NewUser(email, displayName, password string) User {
	id, _ := uuid.NewV7()

	u := User{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}

	u.SetEmail(email)
	u.SetDisplayName(displayName)
	u.SetPassword(password)

	return u
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Email() string {
	return u.email
}

func (u *User) SetEmail(newEmail string) {
	u.email = newEmail
}

func (u User) DisplayName() string {
	return u.displayName
}

func (u *User) SetDisplayName(newDisplayName string) {
	u.displayName = newDisplayName
	u.slug = slugify(newDisplayName)
}

func slugify(input string) string {
	l := strings.ToLower(input)
	return strings.ReplaceAll(l, " ", "-")
}

func (u User) Slug() string {
	return u.slug
}

func (u User) PasswordHash() string {
	return u.password
}

func (u *User) SetPassword(password string) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u.password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(testPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.password), []byte(testPassword))
}

func (u *User) AddRole(newRole UserRole) {
	if u.HasRole(newRole) {
		return
	}
	u.roles = append(u.roles, newRole)
}

func (u User) HasRole(role UserRole) bool {
	for _, userRole := range u.roles {
		if userRole == role {
			return true
		}
	}

	return false
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u User) MarshalJSON() ([]byte, error) {
	dataMap := map[string]any{
		"id":          u.id.String(),
		"email":       u.email,
		"displayName": u.displayName,
		"slug":        u.slug,
		"createdAt":   u.createdAt,
		"updatedAt":   u.updatedAt,
	}

	return json.Marshal(dataMap)
}

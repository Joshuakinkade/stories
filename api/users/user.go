package users 

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/joshuakinkade/web-journal/db"
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

	isNew        bool
	rolesAdded   []uuid.UUID
	rolesRemoved []uuid.UUID
}

func NewUser(email, displayName, password string) User {
	id, _ := uuid.NewV7()

	u := User{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		isNew:     true,
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

func (u *User) Save() error {
	u.updatedAt = time.Now()
	if u.isNew {
		err := u.insert(context.Background())
		if err != nil {
			return err
		}
		u.isNew = false
	} else {
		err := u.update(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}

const insertStmt string = "INSERT INTO users(id, email, display_name, slug, password, roles, created_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8);"

func (u *User) insert(ctx context.Context) error {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return err
	}
	
	_, err = conn.Exec(ctx, insertStmt, u.ID(), u.Email(), u.DisplayName(), u.Slug(), u.PasswordHash(), u.getStringRoles(), u.CreatedAt(), u.UpdatedAt())
	return err
}

func (u *User) getStringRoles() []string {
	var roleStrings []string
	for _, r := range u.roles {
		roleStrings = append(roleStrings, r.String())
	}
	return roleStrings
}

const updateStmt string = "UPDATE users SET email = $2, display_name = $3, slug = $4 password = $5, roles = $6, updated_at = $7 WHERE id = $1"

func (u *User) update(ctx context.Context) error {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return err
	}
	_, err = conn.Exec(ctx, updateStmt, u.ID(), u.Email(), u.DisplayName(), u.Slug(), u.PasswordHash(), u.getStringRoles(), u.UpdatedAt())
	return err
}

func FindByID(id uuid.UUID) (User, error) {
	return User{}, nil
}

const findByEmailStmt string = "SELECT id, email, display_name, slug, password, created_at, updated_at FROM users WHERE email = $1"

func FindByEmail(ctx context.Context, targetEmail string) (User, error) {
	conn, err := db.GetConnection(ctx)
	if err != nil {
		return User{}, err
	}
	row := conn.QueryRow(ctx, findByEmailStmt, targetEmail)
	var id uuid.UUID
	var email string
	var displayName string
	var slug string
	var passwordHash string
	var createdAt time.Time
	var updatedAt time.Time
	err = row.Scan(&id, &email, &displayName, &slug, &passwordHash, &createdAt, &updatedAt)
	if err != nil {
		return User{}, err
	}

	u := User{
		id:          id,
		email:       email,
		displayName: displayName,
		slug:        slug,
		password:    passwordHash,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}

	return u, nil
}

func scanUser() (User, error) {
	return User{}, nil
}

func (u User) MarshalJSON() ([]byte, error) {
	dataMap := map[string]any{
		"id":          u.id.String(),
		"email":       u.email,
		"displayName": u.displayName,
		"createdAt":   u.createdAt,
		"updatedAt":   u.updatedAt,
	}

	return json.Marshal(dataMap)
}

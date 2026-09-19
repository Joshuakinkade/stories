package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateAccessToken(userID uuid.UUID) (string, error) {
	expiration := time.Now().Add(time.Hour * 24)
	claims := jwt.MapClaims{
		"exp": expiration,
		"sub": userID.String(),
	}
	return createTokenWithClaims(claims)
}

func CreateRefreshToken(userID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID.String(),
	}
	return createTokenWithClaims(claims)
}

func createTokenWithClaims(claims jwt.MapClaims) (string, error) {
	key := os.Getenv("AUTH_KEY")
	if key == "" {
		log.Panic("Environment error: AUTH_KEY not set!")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(key))
}

func CheckJWT() error {
	return nil
}

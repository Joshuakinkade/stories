package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Manager struct {
	authkey    []byte
	refreshKey []byte
}

func New(authKey []byte, refreshKey []byte) Manager {
	return Manager{
		authkey:    authKey,
		refreshKey: refreshKey,
	}
}

func (as Manager) CreateAccessToken(userID uuid.UUID) (string, error) {
	expiration := time.Now().Add(time.Hour * 24)
	claims := jwt.MapClaims{
		"exp": expiration.Unix(),
		"sub": userID.String(),
		"iat": time.Now().Unix(),
		"jti": uuid.New(),
	}
	return as.createTokenWithClaims(claims, as.authkey)
}

func (as Manager) CreateRefreshToken(userID uuid.UUID) (string, error) {
	expiration := time.Now().Add(time.Hour * 24 * 7)
	claims := jwt.MapClaims{
		"exp": expiration.Unix(),
		"sub": userID.String(),
		"iat": time.Now().Unix(),
		"jti": uuid.New(),
	}
	return as.createTokenWithClaims(claims, as.refreshKey)
}

func (as Manager) createTokenWithClaims(claims jwt.MapClaims, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func (as Manager) CheckJWT(tokenString string) (string, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (any, error) {
		return as.authkey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("Auth Error")
	}

	userId, err := token.Claims.GetSubject()
	if err != nil {
		return "", errors.New("Could not get userID")
	}

	return userId, nil
}

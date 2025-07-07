package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt       = "qweasdzxc" // хорошая практика добавлять соль к паролю
	signingKey = "qweasadasfasfasfsfafasf"
	// tokenTTL   = 12 * time.Hour
	tokenTTL = 10 * time.Minute
)

type TokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetSigningKey() string {
	return signingKey
}

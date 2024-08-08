package utils

import (
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "qweasdzxc" // хорошая практика добавлять соль к паролю
	signingKey = "qweasadasfasfasfsfafasf"
	// tokenTTL   = 12 * time.Hour
	tokenTTL = 10 * time.Minute
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

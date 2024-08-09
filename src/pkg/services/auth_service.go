package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"practice/domains"
	"practice/pkg/repositories"
	"practice/pkg/utils"
	"time"
)

const (
	salt       = "qweasdzxc" // хорошая практика добавлять соль к паролю
	signingKey = "qweasadasfasfasfsfafasf"
	// tokenTTL   = 12 * time.Hour
	tokenTTL = 10 * time.Minute
)

type AuthService struct {
	auth repositories.Authorization
	user repositories.Users
}

func NewAuthService(
	repo repositories.Authorization,
	user repositories.Users,
) *AuthService {
	return &AuthService{
		auth: repo,
		user: user,
	}
}

func (s *AuthService) CreateUser(user domains.User) (int, error) {
	user.Password = utils.GeneratePasswordHash(user.Password)
	user, err := s.auth.CreateUser(user)

	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.user.GetUserByEmail(email)

	if user.Id == 0 {
		return "User not found", err
	}
	if err != nil {
		return "", err
	}

	inputPasswordHash := utils.GeneratePasswordHash(password)

	if inputPasswordHash != user.Password {
		return "", errors.New("Password is not equal")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &utils.TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	token_created, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	// todo save token in Redis

	return token_created, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &utils.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*utils.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

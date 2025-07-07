package services

import (
	"errors"
	"practice/domains"
	"practice/pkg/repositories"
	"practice/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
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
	hashedPassword, err := utils.GeneratePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword

	createdUser, err := s.auth.CreateUser(user)
	if err != nil {
		return 0, err
	}
	return createdUser.Id, nil
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.user.GetUserByEmail(email)

	if user.Id == 0 {
		return "", errors.New("User not found")
	}
	if err != nil {
		return "", err
	}

	// Проверяем пароль используя bcrypt
	if err := utils.ComparePasswordHash(user.Password, password); err != nil {
		return "", errors.New("Password is not equal")
	}

	// Создаем JWT токен с новой структурой claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &utils.TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: user.Id,
	})

	signingKey := utils.GetSigningKey()
	tokenString, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}

	// todo save token in Redis

	return tokenString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &utils.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(utils.GetSigningKey()), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*utils.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

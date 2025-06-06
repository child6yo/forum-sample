package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/child6yo/forum-sample"
	"github.com/child6yo/forum-sample/pkg/validation"
	"github.com/child6yo/forum-sample/internal/app/repository"
	"github.com/golang-jwt/jwt/v5"
)

var (
	salt = "ojnwsdo342234onjs3"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL = 3 * time.Hour
)

type tokenClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user forum.User) (int, error) {
	v := validation.NewValidator()
	if err := v.ValidateStruct(user); err != nil {
		return 0, err
	}
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	claims := tokenClaims{
		user.Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

package service

import (
	"github.com/child6yo/forum-sample"
	"github.com/child6yo/forum-sample/pkg/repository"
)

type Authorization interface {
	CreateUser(user forum.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
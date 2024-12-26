package repository

import (
	"github.com/child6yo/forum-sample"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user forum.User) (int, error)
	GetUser(username, password string) (forum.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuth(db),
	}
}
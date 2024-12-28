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

type Posts interface {
	CreatePost(post forum.Posts) (forum.Posts, error)
	GetPostById(id int) (forum.Posts, error)
	GetAllPosts() ([]forum.PostsList, error)
	UpdatePost(userId, postId int, input forum.UpdatePostInput) error
	DeletePost(userId, postId int) error
}

type Service struct {
	Authorization
	Posts
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Posts:         NewPostsServise(repos.Posts),
	}
}

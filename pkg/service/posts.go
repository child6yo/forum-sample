package service

import (
	"time"

	"github.com/child6yo/forum-sample"
	"github.com/child6yo/forum-sample/pkg/repository"
)

type PostsServise struct {
	repo repository.Posts
}

func NewPostsServise(repo repository.Posts) *PostsServise {
	return &PostsServise{repo: repo}
}

func (s *PostsServise) CreatePost(post forum.Posts) (forum.Posts, error) {
	time := time.Now()
	post.CrDate = time
	return s.repo.CreatePost(post)
}

func (s *PostsServise) GetById(id int) (forum.Posts, error) {
	return s.repo.GetPostById(id)
}

func (s *PostsServise) GetAllPosts() ([]forum.PostsList, error) {
	return s.repo.GetAllPosts()
}
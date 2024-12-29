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

func (s *PostsServise) CreatePost(post forum.Posts) (int, error) {
	time := time.Now()
	post.CrTime = time
	return s.repo.CreatePost(post)
}

func (s *PostsServise) GetPostById(id int) (forum.Posts, error) {
	return s.repo.GetPostById(id)
}

func (s *PostsServise) GetAllPosts() ([]forum.PostsList, error) {
	return s.repo.GetAllPosts()
}

func (s *PostsServise) UpdatePost(userId, postId int, input forum.UpdatePostInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	time := time.Now()
	input.UpdTime = &time
	return s.repo.UpdatePost(userId, postId, input)
}

func (s *PostsServise) DeletePost(userId, postId int) error {
	return s.repo.DeletePost(userId, postId)
}

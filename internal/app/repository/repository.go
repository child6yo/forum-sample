package repository

import (
	"github.com/child6yo/forum-sample"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user forum.User) (int, error)
	GetUser(username, password string) (forum.User, error)
}

type Posts interface {
	CreatePost(post forum.Posts) (int, error)
	GetPostById(id int) (forum.Posts, error)
	GetAllPosts() ([]forum.PostsList, error)
	UpdatePost(userId, postId int, input forum.UpdatePostInput) error
	DeletePost(userId, postId int) error
}

type Threads interface {
	ThreadExists(threadId, postId int) (bool, error)
	CreateThread(postId int, thread forum.Threads) (int, error)
	GetThreadById(threadId int) (forum.Threads, error)
	GetThreadsByPost(postId int) ([]forum.Threads, error)
	UpdateThread(userId, threadId int, input forum.UpdateThreadInput) error
}

type Repository struct {
	Authorization
	Posts
	Threads
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthDatabase(db),
		Posts:         NewPostsDatabase(db),
		Threads:       NewThreadsDatabase(db),
	}
}
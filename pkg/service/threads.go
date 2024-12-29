package service

import (
	"time"

	"github.com/child6yo/forum-sample"
	"github.com/child6yo/forum-sample/pkg/repository"
)

type ThreadsServise struct {
	repo repository.Threads
}

func NewThreadsServise(repo repository.Threads) *ThreadsServise {
	return &ThreadsServise{repo: repo}
}

func (s *ThreadsServise) CreateThread(postId int, thread forum.Threads) (int, error) {
	time := time.Now()
	thread.CrTime = time
	return s.repo.CreateThread(postId, thread)
}

func (s *ThreadsServise)GetThreadById(threadId int) (forum.Threads, error) {
	return s.repo.GetThreadById(threadId)
}

func (s *ThreadsServise)GetThreadsByPost() {

}

func (s *ThreadsServise)UpdateThread() {

}

func (s *ThreadsServise)DeleteThread() {
	
}
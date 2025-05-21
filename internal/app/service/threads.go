package service

import (
	"time"

	"github.com/child6yo/forum-sample"
	"github.com/child6yo/forum-sample/internal/app/repository"
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

func (s *ThreadsServise) GetThreadById(threadId int) (forum.Threads, error) {
	return s.repo.GetThreadById(threadId)
}

func (s *ThreadsServise) GetThreadsByPost(postId int) ([]*forum.ThreadsList, error) {
	threads, err := s.repo.GetThreadsByPost(postId)
	if err != nil {
		return []*forum.ThreadsList{}, err
	}
	orgThr := OrganizeThreads(threads)

	return orgThr, nil
}

func (s *ThreadsServise) UpdateThread(userId, threadId int, input forum.UpdateThreadInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	time := time.Now()
	input.UpdTime = &time
	return s.repo.UpdateThread(userId, threadId, input)
}

func OrganizeThreads(threads []forum.Threads) []*forum.ThreadsList {
	idToThList := make(map[int]*forum.ThreadsList)
	var ThList []*forum.ThreadsList

	for _, th := range threads {
		idToThList[th.Id] = &forum.ThreadsList{
			Id:   th.Id,
			UserId: th.UserId,
			Content: th.Content,
			CrTime: th.CrTime,
			Update: th.Update,
			UpdTime: th.UpdTime,
			Answers:  []*forum.ThreadsList{},
		}
		if th.AnswerAt == 0 {
			ThList = append(ThList, idToThList[th.Id])
		} else {
			parent := idToThList[th.AnswerAt]
			parent.Answers = append(parent.Answers, idToThList[th.Id])
		}
	}
	
	return ThList
}
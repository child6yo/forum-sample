package forum

import "time"

type Threads struct {
	Id       int       `json:"id" db:"id"`
	UserId   int       `json:"user_id" db:"user_id"`
	Content  string    `json:"content" db:"content" binding:"required"`
	AnswerAt int       `json:"answer_at" db:"answer_at"`
	CrTime   time.Time `json:"cr_time" db:"cr_time"`
	Update   bool      `json:"update" db:"update"`
	UpdTime  time.Time `json:"upd_time" db:"upd_time"`
}

type PostThreads struct {
	Id       int `json:"id" db:"id"`
	PostId   int `json:"post_id" db:"post_id"`
	ThreadId int `json:"thread_id" db:"thread_id"`
}

type ThreadsList struct {
	Id      int       `json:"id"`
	UserId  int       `json:"user_id"`
	Content string    `json:"content"`
	CrTime  time.Time `json:"cr_time"`
	Update  bool      `json:"update"`
	UpdTime time.Time `json:"upd_time"`
	Answers []*ThreadsList `json:"threads"`
}

package forum

import (
	"errors"
	"time"
)

type Posts struct {
	Id      int       `json:"id" db:"id"`
	UserId  int       `json:"user_id" db:"user_id"`
	Title   string    `json:"title" db:"title" binding:"required"`
	Content string    `json:"content" db:"content" binding:"required"`
	CrTime  time.Time `json:"cr_time" db:"cr_time"`
	Update  bool      `json:"update" db:"update"`
	UpdTime time.Time `json:"upd_time" db:"upd_time"`
}

type PostsList struct {
	Id      int       `json:"id" db:"id"`
	UserId  int       `json:"user_id" db:"user_id"`
	Title   string    `json:"title" db:"title"`
	CrTime  time.Time `json:"cr_time" db:"cr_time"`
	UpdTime time.Time `json:"upd_time" db:"upd_time"`
}

type UpdatePostInput struct {
	Title   *string    `json:"title"`
	Content *string    `json:"content"`
	UpdTime *time.Time `json:"upd_time"`
}

func (i UpdatePostInput) Validate() error {
	if i.Title == nil && i.Content == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

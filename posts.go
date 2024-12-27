package forum

import "time"

type Posts struct {
	Id      int       `json:"id" db:"id"`
	UserId  int       `json:"user_id" db:"user_id"`
	Title   string    `json:"title" db:"title" binding:"required"`
	Content string    `json:"content" db:"content" binding:"required"`
	CrDate  time.Time `json:"cr_date" db:"cr_date"`
}

type PostsList struct {
	Id      int       `json:"id" db:"id"`
	UserId  int       `json:"user_id" db:"user_id"`
	Title  string    `json:"title" db:"title"`
	CrDate time.Time `json:"cr_date" db:"cr_date"`
}

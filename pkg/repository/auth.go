package repository

import (
	"fmt"

	"github.com/child6yo/forum-sample"
	"github.com/jmoiron/sqlx"
)

type AuthDatabase struct {
	db *sqlx.DB
}

func NewAuthDatabase(db *sqlx.DB) *AuthDatabase {
	return &AuthDatabase{db: db}
}

func (r *AuthDatabase) CreateUser(user forum.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Email, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthDatabase) GetUser(username, password string) (forum.User, error) {
	var user forum.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
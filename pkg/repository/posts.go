package repository

import (
	"fmt"
	"strings"

	"github.com/child6yo/forum-sample"
	"github.com/jmoiron/sqlx"
)

type PostsDatabase struct {
	db *sqlx.DB
}

func NewPostsDatabase(db *sqlx.DB) *PostsDatabase {
	return &PostsDatabase{db: db}
}

func (r *PostsDatabase) CreatePost(post forum.Posts) (int, error) {
	var id int

	query := fmt.Sprintf(`INSERT INTO %s (user_id, title, content, cr_time, update, upd_time)
						 values ($1, $2, $3, $4, $5, $6) RETURNING id`, postsTable)

	row := r.db.QueryRow(query, post.UserId, post.Title, post.Content, post.CrTime, post.Update, post.UpdTime)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *PostsDatabase) GetPostById(postId int) (forum.Posts, error) {
	var post forum.Posts

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postsTable)
	err := r.db.Get(&post, query, postId)

	return post, err
}

func (r *PostsDatabase) GetAllPosts() ([]forum.PostsList, error) {
	var posts []forum.PostsList

	query := fmt.Sprintf("SELECT id, user_id, title, cr_time, upd_time FROM %s", postsTable)
	err := r.db.Select(&posts, query)

	return posts, err
}

func (r *PostsDatabase) UpdatePost(userId, postId int, input forum.UpdatePostInput) error {
	var id int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE id=$1", postsTable)
	if err := r.db.Get(&id, query, postId); err != nil {
		return err
	} else if id != userId {
		return fmt.Errorf("attempt to update someone else's post")
	}

	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Content != nil {
		setValues = append(setValues, fmt.Sprintf("content=$%d", argId))
		args = append(args, *input.Content)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("upd_time=$%d", argId))
	args = append(args, *input.UpdTime)
	argId++

	setQuery := strings.Join(setValues, ", ")

	query = fmt.Sprintf("UPDATE %s SET %s, update=true WHERE id=$%d",
		postsTable, setQuery, argId)
	args = append(args, postId)
	_, err := r.db.Exec(query, args...)
	return err
}

func (r *PostsDatabase) DeletePost(userId, postId int) error {
	var id int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE id=$1", postsTable)
	if err := r.db.Get(&id, query, postId); err != nil {
		return err
	} else if id != userId {
		return fmt.Errorf("attempt to delete someone else's post")
	}

	tx, err := r.db.Begin()
    if err != nil {
        return err
    }

	postThreadsQuery := fmt.Sprintf("DELETE FROM %s t USING %s pt WHERE pt.post_id=$1 AND t.id=pt.thread_id", threadsTable, postThreadsTable)
	_, err = tx.Exec(postThreadsQuery, postId)
    if err != nil {
        tx.Rollback()
        return err
    }

	postQuery := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postsTable)
	_, err = tx.Exec(postQuery, postId)
	if err != nil {
        tx.Rollback()
        return err
    }

	return tx.Commit()
}
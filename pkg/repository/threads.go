package repository

import (
	"fmt"

	"github.com/child6yo/forum-sample"
	"github.com/jmoiron/sqlx"
)

type ThreadsDatabase struct {
	db *sqlx.DB
}

func NewThreadsDatabase(db *sqlx.DB) *ThreadsDatabase {
	return &ThreadsDatabase{db: db}
}

func (r *ThreadsDatabase) PostExists(postID int) (bool, error) {
    var exists bool
    query := "SELECT EXISTS(SELECT 1 FROM threads WHERE id=$1)"
    err := r.db.QueryRow(query, postID).Scan(&exists)
    return exists, err
}

func (r *ThreadsDatabase) CreateThread(postId int, thread forum.Threads) (int, error) {
    tx, err := r.db.Begin()
    if err != nil {
        return 0, err
    }

    if thread.AnswerAt != 0 {
        exists, err := r.PostExists(thread.AnswerAt)
        if err != nil {
            tx.Rollback()
            return 0, err
        }
        if !exists {
            tx.Rollback()
            return 0, fmt.Errorf("post with ID %d does not exist", thread.AnswerAt)
        }
    }

    var threadId int
    threadQuery := fmt.Sprintf(`INSERT INTO %s (user_id, content, answer_at, cr_time, update, upd_time) 
                                values ($1, $2, $3, $4, $5, $6) RETURNING id`, threadsTable)
    row := tx.QueryRow(threadQuery, thread.UserId, thread.Content, thread.AnswerAt, thread.CrTime, thread.Update, thread.UpdTime)
    err = row.Scan(&threadId)
    if err != nil {
        tx.Rollback()
        return 0, err
    }

    createListItemsQuery := fmt.Sprintf("INSERT INTO %s (post_id, thread_id) values ($1, $2)", postThreadsTable)
    _, err = tx.Exec(createListItemsQuery, postId, threadId)
    if err != nil {
        tx.Rollback()
        return 0, err
    }

    return threadId, tx.Commit()
}

func (r *ThreadsDatabase) GetThreadById(threadId int) (forum.Threads, error) {
	var thread forum.Threads

	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", threadsTable)
	err := r.db.Get(&thread, query, threadId)

	return thread, err
}

func (r *ThreadsDatabase) GetThreadTree() {

}

func (r *ThreadsDatabase) GetThreadsByPost() {

}

func (r *ThreadsDatabase) UpdateThread() {

}

func (r *ThreadsDatabase) DeleteThread() {
	
}
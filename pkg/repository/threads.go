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

func (r *ThreadsDatabase) ThreadExists(threadId, postId int) (bool, error) {
    var exists bool
    query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s t 
                            INNER JOIN %s pt ON t.id = pt.thread_id WHERE t.id=$1 AND pt.post_id=$2)`, threadsTable, postThreadsTable)
    err := r.db.QueryRow(query, threadId, postId).Scan(&exists)
    return exists, err
}

func (r *ThreadsDatabase) CreateThread(postId int, thread forum.Threads) (int, error) {
    tx, err := r.db.Begin()
    if err != nil {
        return 0, err
    }

    if thread.AnswerAt != 0 {
        exists, err := r.ThreadExists(thread.AnswerAt, postId)
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

func (r *ThreadsDatabase) GetThreadsByPost(postId int) ([]forum.Threads, error) {
    var threads []forum.Threads

    query := fmt.Sprintf(`SELECT t.id, t.user_id, t.content, t.answer_at, t.cr_time, t.update, t.upd_time 
                            FROM %s t INNER JOIN %s pt on pt.thread_id = t.id WHERE pt.post_id=$1
                            ORDER BY t.id`, threadsTable, postThreadsTable)
	err := r.db.Select(&threads, query, postId)

	return threads, err
}

func (r *ThreadsDatabase) UpdateThread(userId, threadId int, input forum.UpdateThreadInput) error {
	var id int
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE id=$1", threadsTable)
	if err := r.db.Get(&id, query, threadId); err != nil {
		return err
	} else if id != userId {
		return fmt.Errorf("not your thread)")
	}

	query = fmt.Sprintf("UPDATE %s SET content=$1, update=true, upd_time=$2 WHERE id=%d",
		threadsTable, threadId)
	_, err := r.db.Exec(query, input.Content, input.UpdTime)
	return err
}
package repository

import (
	"fmt"

	"github.com/child6yo/forum-sample"
	"github.com/jmoiron/sqlx"
)

type PostsDatabase struct {
	db *sqlx.DB
}

func NewPostsDatabase(db *sqlx.DB) *PostsDatabase {
	return &PostsDatabase{db: db}
}

func (r *PostsDatabase) CreatePost(post forum.Posts) (forum.Posts, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (user_id, title, content, cr_date) values ($1, $2, $3, $4) RETURNING id", postsTable)

	row := r.db.QueryRow(query, post.UserId, post.Title, post.Content, post.CrDate)
	if err := row.Scan(&id); err != nil {
		return forum.Posts{}, err
	}

	newPost := forum.Posts{Id: id, UserId: post.UserId, Title: post.Title, Content: post.Content, CrDate: post.CrDate}
	return newPost, nil
}

func (r *PostsDatabase) GetPostById(postId int) (forum.Posts, error) {
	var post forum.Posts
	
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postsTable)
	err := r.db.Get(&post, query, postId)

	return post, err
}

func (r *PostsDatabase) GetAllPosts() ([]forum.PostsList, error) {
	var posts []forum.PostsList
	
	query := fmt.Sprintf("SELECT id, user_id, title, cr_date FROM %s", postsTable)
	err := r.db.Select(&posts, query)

	return posts, err
}
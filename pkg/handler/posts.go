package handler

import (
	"net/http"
	"strconv"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, "create post", http.StatusForbidden, err)
		return
	}

	var input forum.Posts
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "create post", http.StatusBadRequest, err)
		return
	}
	input.UserId = userId

	id, err := h.services.Posts.CreatePost(input)
	if err != nil {
		errorResponse(c, "create post", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "create post", map[string]interface{}{
		"post id": id,
	})
}

func (h *Handler) getPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, "get post by id", http.StatusForbidden, err)
		return
	}

	post, err := h.services.Posts.GetPostById(id)
	if err != nil {
		errorResponse(c, "get post by id", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "create post", post)
}

func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.Posts.GetAllPosts()
	if err != nil {
		errorResponse(c, "get all posts", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "get all posts", posts)
}

func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, "update post", http.StatusForbidden, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, "update post", http.StatusBadRequest, err)
		return
	}

	var input forum.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "update post", http.StatusBadRequest, err)
		return
	}

	if err := h.services.Posts.UpdatePost(userId, id, input); err != nil {
		errorResponse(c, "update post", http.StatusForbidden, err)
		return
	}

	successResponse(c, "update post", nil)
}

func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, "delete post", http.StatusForbidden, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, "delete post", http.StatusBadRequest, err)
		return
	}

	if err := h.services.Posts.DeletePost(userId, id); err != nil {
		errorResponse(c, "delete post", http.StatusForbidden, err)
		return
	}

	successResponse(c, "delete post", nil)
}
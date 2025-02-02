package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

// @Summary Create post
// @Security ApiKeyAuth
// @Tags Posts
// @Description create post
// @ID create-post
// @Accept  json
// @Produce  json
// @Param input body forum.Posts true "post info"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/posts [post]
func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		err = fmt.Errorf("unknown user")
		errorResponse(c, "create post", http.StatusForbidden, err)
		return
	}

	var input forum.Posts
	if err := c.BindJSON(&input); err != nil {
		err = fmt.Errorf("invalid request body")
		errorResponse(c, "create post", http.StatusBadRequest, err)
		return
	}
	input.UserId = userId

	id, err := h.services.Posts.CreatePost(input)
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "create post", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "create post", map[string]interface{}{
		"post_id": id,
	})
}

// @Summary Get post
// @Security ApiKeyAuth
// @Tags Posts
// @Description get post by id
// @ID get-post
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/posts/:id [get]
func (h *Handler) getPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "get post by id", http.StatusForbidden, err)
		return
	}

	post, err := h.services.Posts.GetPostById(id)
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "get post by id", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "create post", post)
}

// @Summary Get all posts
// @Security ApiKeyAuth
// @Tags Posts
// @Description get all posts 
// @ID get-all-posts
// @Accept  json
// @Produce  json
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/posts [get]
func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.Posts.GetAllPosts()
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "get all posts", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "get all posts", posts)
}

// @Summary Update post
// @Security ApiKeyAuth
// @Tags Posts
// @Description update post 
// @ID update-post
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param input body forum.Posts true "post info"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/posts/:id [put]
func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		err = fmt.Errorf("unknown user")
		errorResponse(c, "update post", http.StatusForbidden, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "update post", http.StatusBadRequest, err)
		return
	}

	var input forum.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		err = fmt.Errorf("invalid request body")
		errorResponse(c, "update post", http.StatusBadRequest, err)
		return
	}

	if err := h.services.Posts.UpdatePost(userId, id, input); err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "update post", http.StatusForbidden, err)
		return
	}

	successResponse(c, "update post", nil)
}

// @Summary Delete post
// @Security ApiKeyAuth
// @Tags Posts
// @Description delete post 
// @ID delete-post
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/posts/:id [delete]
func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		err = fmt.Errorf("unknown user")
		errorResponse(c, "delete post", http.StatusForbidden, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "delete post", http.StatusBadRequest, err)
		return
	}

	if err := h.services.Posts.DeletePost(userId, id); err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "delete post", http.StatusForbidden, err)
		return
	}

	successResponse(c, "delete post", nil)
}

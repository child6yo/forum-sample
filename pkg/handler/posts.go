package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input forum.Posts
	if err := c.BindJSON(&input); err != nil {
		log.Fatal("ne bindit")
		return
	}
	input.UserId = userId

	post, err := h.services.Posts.CreatePost(input)
	if err != nil {
		log.Fatal("ne creatit ", err)
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) getPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	post, err := h.services.Posts.GetPostById(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, post)
}

type allPosts struct {
	Posts []forum.PostsList `json:"posts"`
}

func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.Posts.GetAllPosts()
	if err != nil {
		log.Fatal("ne daet", err)
		return
	}

	c.JSON(http.StatusOK, allPosts{Posts: posts})
}

func (h *Handler) updatePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	var input forum.UpdatePostInput
	if err := c.BindJSON(&input); err != nil {
		return
	}

	if err := h.services.Posts.UpdatePost(userId, id, input); err != nil {
		log.Print("err: ", err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (h *Handler) deletePost(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	if err := h.services.Posts.DeletePost(userId, id); err != nil {
		log.Print("err: ", err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
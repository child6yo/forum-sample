package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost(c *gin.Context) {
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
		log.Fatal("ne creatit")
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *Handler) GetPostById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	post, err := h.services.Posts.GetById(id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, post)
}

type allPosts struct {
	Posts []forum.PostsList `json:"posts"`
}

func (h *Handler) GetAllPosts(c *gin.Context) {
	posts, err := h.services.Posts.GetAllPosts()
	if err != nil {
		log.Fatal("ne daet", err)
		return
	}

	c.JSON(http.StatusOK, allPosts{Posts: posts})
}
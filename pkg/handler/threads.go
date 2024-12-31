package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateThread(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		log.Fatal(err)
		return
	}
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("? че ", err)
		return
	}
	answerAt, err := strconv.Atoi(c.DefaultQuery("answer", "0"))
	if err != nil {
		log.Fatal("?? ", err)
		return
	}

	var input forum.Threads
	if err := c.BindJSON(&input); err != nil {
		log.Fatal(err)
		return
	}
	input.UserId = userId
	input.AnswerAt = answerAt

	thread, err := h.services.Threads.CreateThread(postId, input)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, thread)
}

func (h *Handler) GetThreadById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("?")
		return
	}

	thread, err := h.services.Threads.GetThreadById(id)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, thread)
}

func (h *Handler) GetThreadByPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Fatal("?")
		return
	}

	threads, err := h.services.Threads.GetThreadsByPost(id)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, threads)
}
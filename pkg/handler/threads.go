package handler

import (
	"net/http"
	"strconv"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateThread(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, "create thread", http.StatusForbidden, err)
		return
	}
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, "create thread", http.StatusBadRequest, err)
		return
	}
	answerAt, err := strconv.Atoi(c.DefaultQuery("answer", "0"))
	if err != nil {
		errorResponse(c, "create thread", http.StatusBadRequest, err)
		return
	}

	var input forum.Threads
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "create thread", http.StatusBadRequest, err)
		return
	}
	input.UserId = userId
	input.AnswerAt = answerAt

	id, err := h.services.Threads.CreateThread(postId, input)
	if err != nil {
		errorResponse(c, "create thread", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "create thread", map[string]interface{}{
		"thread id": id,
	})
}

func (h *Handler) GetThreadById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, "get thread by id", http.StatusBadRequest, err)
		return
	}

	thread, err := h.services.Threads.GetThreadById(id)
	if err != nil {
		errorResponse(c, "get thread by id", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "get thread by id", thread)
}

func (h *Handler) GetThreadByPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, "get thread by post", http.StatusBadRequest, err)
		return
	}

	threads, err := h.services.Threads.GetThreadsByPost(id)
	if err != nil {
		errorResponse(c, "get thread by post", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "get thread by post", threads)
}

func (h *Handler) UpdateThread(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		errorResponse(c, "update thread", http.StatusForbidden, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, "update thread", http.StatusBadRequest, err)
		return
	}

	var input forum.UpdateThreadInput
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "update thread", http.StatusBadRequest, err)
		return
	}

	if err := h.services.Threads.UpdateThread(userId, id, input); err != nil {
		errorResponse(c, "update thread", http.StatusForbidden, err)
		return
	}

	successResponse(c, "update thread", nil)
}
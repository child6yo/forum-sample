package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

// @Summary Create thread
// @Security ApiKeyAuth
// @Tags Threads
// @Description create thread
// @ID create-thread
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param answer query string false "answer at (another thread). deault=0"
// @Param input body forum.Threads true "thread info"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/posts/:id/threads [post]
func (h *Handler) CreateThread(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		err = fmt.Errorf("unknown user")
		errorResponse(c, "create thread", http.StatusForbidden, err)
		return
	}
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "create thread", http.StatusBadRequest, err)
		return
	}
	answerAt, err := strconv.Atoi(c.DefaultQuery("answer", "0"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "create thread", http.StatusBadRequest, err)
		return
	}

	var input forum.Threads
	if err := c.BindJSON(&input); err != nil {
		err = fmt.Errorf("invalid request body")
		errorResponse(c, "create thread", http.StatusBadRequest, err)
		return
	}
	input.UserId = userId
	input.AnswerAt = answerAt

	id, err := h.services.Threads.CreateThread(postId, input)
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "create thread", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "create thread", map[string]interface{}{
		"thread id": id,
	})
}

// @Summary Get thread
// @Security ApiKeyAuth
// @Tags Threads
// @Description get thread by id
// @ID get-thread
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/thread/:id [get]
func (h *Handler) GetThreadById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "get thread by id", http.StatusBadRequest, err)
		return
	}

	thread, err := h.services.Threads.GetThreadById(id)
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "get thread by id", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "get thread by id", thread)
}

// @Summary Get threads
// @Security ApiKeyAuth
// @Tags Threads
// @Description get threads by post
// @ID get-threads
// @Accept  json
// @Produce  json
// @Param id path int true "post id"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/posts/:id/threads/ [get]
func (h *Handler) GetThreadByPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "get thread by post", http.StatusBadRequest, err)
		return
	}

	threads, err := h.services.Threads.GetThreadsByPost(id)
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "get thread by post", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "get thread by post", threads)
}

// @Summary Update thread
// @Security ApiKeyAuth
// @Tags Threads
// @Description update thread by id
// @ID update-thread
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} Response
// @Failure 400,404 {object} Response
// @Failure 500 {object} Response
// @Failure default {object} Response
// @Router /api/v1/thread/:id [put]
func (h *Handler) UpdateThread(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		err = fmt.Errorf("unknown user")
		errorResponse(c, "update thread", http.StatusForbidden, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		err = fmt.Errorf("invalid request")
		errorResponse(c, "update thread", http.StatusBadRequest, err)
		return
	}

	var input forum.UpdateThreadInput
	if err := c.BindJSON(&input); err != nil {
		err = fmt.Errorf("invalid request body")
		errorResponse(c, "update thread", http.StatusBadRequest, err)
		return
	}

	if err := h.services.Threads.UpdateThread(userId, id, input); err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "update thread", http.StatusForbidden, err)
		return
	}

	successResponse(c, "update thread", nil)
}
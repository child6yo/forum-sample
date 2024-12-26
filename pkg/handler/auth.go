package handler

import (
	"net/http"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input forum.User

	if err := c.BindJSON(&input); err != nil {
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}


func (h *Handler) signIn(c *gin.Context) {
	var input forum.SignIn

	if err := c.BindJSON(&input); err != nil {
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
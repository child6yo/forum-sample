package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/child6yo/forum-sample"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input forum.User

	if err := c.BindJSON(&input); err != nil {
		err = fmt.Errorf("invalid request data")
		errorResponse(c, "sign up", http.StatusBadRequest, err)
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		err = fmt.Errorf("server error")
		errorResponse(c, "sign up", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "sign up", map[string]interface{}{
		"user_id": id,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input forum.SignIn

	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, "sign in", http.StatusBadRequest, err)
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		errorResponse(c, "sign in", http.StatusInternalServerError, err)
		return
	}

	successResponse(c, "sign in", map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		err := fmt.Errorf("empty authorization header")
		errorResponse(c, "authorization", http.StatusUnauthorized, err)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		err := fmt.Errorf("invalid authorizaton header")
		errorResponse(c, "authorization", http.StatusUnauthorized, err)
		return
	}

	if len(headerParts[1]) == 0 {
		err := fmt.Errorf("token is empty")
		errorResponse(c, "authorization", http.StatusUnauthorized, err)
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		err = fmt.Errorf("invalid token")
		errorResponse(c, "authorization", http.StatusUnauthorized, err)
	}

	c.Set("userId", userId)
}

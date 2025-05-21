package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("unknown jwt")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("unknown jwt")
	}

	return idInt, nil
}


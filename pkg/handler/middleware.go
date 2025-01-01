package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
)

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
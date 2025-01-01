package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int `json:"status"`
	Data   any `json:"data"`
}


func errorResponse(c *gin.Context, description string, httpError int, e error) {
	slog.Error(description, "Error", e)
	c.AbortWithStatusJSON(httpError, Response{Status: httpError, Data: nil})
}

func successResponse(c *gin.Context, description string, data any) {
	slog.Info(description)
	c.JSON(http.StatusOK, Response{Status: http.StatusOK, Data: data})
}
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/child6yo/forum-sample/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		v1 := api.Group("/v1")
		{
			posts := v1.Group("/posts")
			{
				posts.POST("/", h.CreatePost)
				posts.GET("/", h.GetAllPosts)
				posts.GET("/:id", h.GetPostById)
			}
		}
	}

	return router
}
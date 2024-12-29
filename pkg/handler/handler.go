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
				posts.POST("/", h.createPost)
				posts.GET("/", h.getAllPosts)
				posts.GET("/:id", h.getPostById)
				posts.PUT("/:id", h.updatePost)
				posts.DELETE("/:id", h.deletePost)

				threads := posts.Group("/:id/threads")
				{
					threads.POST("/", h.CreateThread)
				}
			}
			threads := posts.Group("/threads")
			{
				threads.GET("/:id", h.GetThreadById)
			}
		}
	}

	return router
}
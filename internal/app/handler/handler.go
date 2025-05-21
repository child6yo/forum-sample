package handler

import (
	"github.com/child6yo/forum-sample/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"

	_ "github.com/child6yo/forum-sample/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
					threads.GET("/", h.GetThreadByPost)
				}
			}
			threads := v1.Group("/threads")
			{
				threads.GET("/:id", h.GetThreadById)
				threads.PUT("/:id", h.UpdateThread)
			}
		}
	}

	return router
}
package handler

import (
	"api/pkg/service"
	"github.com/gin-gonic/gin"
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
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
	}

	//api := router.Group("/api")
	//{
	//	characters := api.Group("/characters")
	//	{
	//		characters.POST("/")
	//		characters.GET("/")
	//		characters.GET("/:id")
	//		characters.PUT("/:id")
	//		characters.DELETE("/:id")
	//	}
	//}

	return router
}

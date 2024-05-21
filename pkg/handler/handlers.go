package handler

import (
	service "rate/pkg/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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

	{
		ratings := router.Group("/ratings")
		{
			ratings.POST("/", )
			ratings.GET("/", h.getActual)
			ratings.PUT("/:id", h.updateRating)
			ratings.DELETE("/:id", h.deleteRating)
	}
}
	return router
}
package handler

import (
	service "rate/pkg/service/mocks"

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
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("/home/r_rmarsu/Rate/assets/index.html")

	rates := router.Group("/rates")
	{
		rates.GET("/main", h.mainweb)
		rates.POST("/rate", h.updateRating)
	}
	return router
}

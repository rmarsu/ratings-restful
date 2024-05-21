package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateRating(c *gin.Context) {

}
func (h *Handler) mainweb(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Оценивалки",
	})
}

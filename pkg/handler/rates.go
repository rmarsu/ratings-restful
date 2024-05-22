package handler

import (
	"net/http"
	rate "rate"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddRating(c *gin.Context) {
	var input rate.Stars
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Stars.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}

func (h *Handler) mainweb(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "Оценивалки",
	})
}

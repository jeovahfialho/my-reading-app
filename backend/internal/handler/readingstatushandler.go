package handler

import (
	"net/http"
	"strconv"

	"my-reading-app/internal/service"

	"github.com/gin-gonic/gin"
)

type ReadingStatusHandler struct {
	service *service.ReadingStatusService
}

func NewReadingStatusHandler(service *service.ReadingStatusService) *ReadingStatusHandler {
	return &ReadingStatusHandler{service: service}
}

func (h *ReadingStatusHandler) GetStatus(c *gin.Context) {
	userId := c.Param("userId")
	statuses, err := h.service.GetStatus(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, statuses)
}

func (h *ReadingStatusHandler) UpdateStatus(c *gin.Context) {
	userId := c.Param("userId")
	day, err := strconv.Atoi(c.Param("day"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day"})
		return
	}

	var req struct {
		Status string `json:"status"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateStatus(userId, day, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}

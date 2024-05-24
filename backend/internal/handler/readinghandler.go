package handler

import (
	"my-reading-app/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReadingHandler struct {
	service service.ReadingService
}

func NewReadingHandler(service service.ReadingService) *ReadingHandler {
	return &ReadingHandler{service: service}
}

func (h *ReadingHandler) GetReading(c *gin.Context) {
	dayStr := c.Param("day")
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day parameter"})
		return
	}

	reading, err := h.service.GetReadingByDay(day)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reading not found"})
		return
	}

	c.JSON(http.StatusOK, reading)
}

func (h *ReadingHandler) NextReading(c *gin.Context) {
	// Implement logic to fetch the next day's reading
}

func (h *ReadingHandler) PreviousReading(c *gin.Context) {
	// Implement logic to fetch the previous day's reading
}

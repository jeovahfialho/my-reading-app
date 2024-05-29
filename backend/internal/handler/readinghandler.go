package handler

import (
	"my-reading-app/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReadingHandler struct {
	readingService service.ReadingService
	bibleService   service.BibleService
}

func NewReadingHandler(readingService service.ReadingService, bibleService service.BibleService) *ReadingHandler {
	return &ReadingHandler{readingService: readingService, bibleService: bibleService}
}

func (h *ReadingHandler) GetReading(c *gin.Context) {
	day := c.Param("day")
	reading, err := h.readingService.GetReading(day)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reading not found"})
		return
	}
	c.JSON(http.StatusOK, reading)
}

func (h *ReadingHandler) GetReadingText(c *gin.Context) {
	description := c.Query("description")
	text, err := h.bibleService.GetBibleText(description)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Text not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"text": text})
}

func (h *ReadingHandler) NextReading(c *gin.Context) {
	dayStr := c.Param("day")
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day parameter"})
		return
	}

	nextDay := day + 1
	reading, err := h.readingService.GetReading(strconv.Itoa(nextDay))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reading not found"})
		return
	}

	c.JSON(http.StatusOK, reading)
}

func (h *ReadingHandler) PreviousReading(c *gin.Context) {
	dayStr := c.Param("day")
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid day parameter"})
		return
	}

	prevDay := day - 1
	if prevDay < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No previous reading"})
		return
	}

	reading, err := h.readingService.GetReading(strconv.Itoa(prevDay))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reading not found"})
		return
	}

	c.JSON(http.StatusOK, reading)
}

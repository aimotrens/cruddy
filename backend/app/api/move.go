package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Move(c *gin.Context) {
	reqPath := c.Query("path")
	reqTarget := c.Query("target")

	fullPath, exists := h.getFullpathChecked(reqPath)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Datei nicht gefunden"})
		return
	}

	targetPath, exists := h.getFullpathChecked(reqTarget)
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Ziel existiert bereits"})
		return
	}

	// Datei verschieben
	err := os.Rename(fullPath, targetPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datei erfolgreich verschoben"})
}

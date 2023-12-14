package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Delete(c *gin.Context) {
	reqPath := c.Query("path")

	fullPath, exists := h.getFullpathChecked(reqPath)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Element nicht gefunden"})
		return
	}

	// Datei löschen
	err := os.RemoveAll(fullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Element erfolgreich gelöscht"})
}

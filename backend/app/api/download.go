package api

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Download(c *gin.Context) {
	reqPath := c.Query("path")

	fullPath, exists := h.getFullpathChecked(reqPath)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Element nicht gefunden"})
		return
	}

	stat, _ := os.Stat(fullPath)
	if stat.IsDir() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Element ist ein Ordner"})
		return
	} else {
		c.Header("Content-Description", "File Transfer")
		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", path.Base(fullPath)))
		c.Header("Content-Type", "application/octet-stream")
		c.File(fullPath)
	}
}

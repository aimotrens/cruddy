package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Mkdir(c *gin.Context) {
	reqPath := c.Query("path")

	fullPath, _ := h.getFullpathChecked(reqPath)

	err := os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ordner erfolgreich erstellt"})
}

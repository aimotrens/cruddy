package api

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Upload(c *gin.Context) {
	reqPath := c.Query("path")

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Datei speichern
	fullPath, exists := h.getFullpathChecked(reqPath)
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "Datei existiert bereits"})
		return
	}

	os.MkdirAll(path.Dir(fullPath), os.ModePerm)

	out, err := os.Create(fullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Datei erfolgreich hochgeladen"})
}

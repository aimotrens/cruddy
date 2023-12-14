package api

import (
	"net/http"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

type ListItem struct {
	FilePath string    `json:"filePath"`
	Name     string    `json:"name"`
	IsDir    bool      `json:"isDir"`
	Size     int64     `json:"size"`
	Changed  time.Time `json:"changed"`
}

func (h *Handler) List(c *gin.Context) {
	reqPath := c.Query("path")

	fullPath, exists := h.getFullpathChecked(reqPath)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pfad nicht gefunden"})
		return
	}

	entries, err := os.ReadDir(fullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := make([]ListItem, 0)
	for _, entry := range entries {
		stat, _ := entry.Info()

		listItem := ListItem{
			FilePath: path.Join(reqPath, entry.Name()),
			Name:     entry.Name(),
			IsDir:    stat.IsDir(),
			Size:     stat.Size(),
			Changed:  stat.ModTime(),
		}

		items = append(items, listItem)
	}

	c.JSON(http.StatusOK, items)
}

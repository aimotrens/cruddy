package api

import (
	"errors"
	"io/fs"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NoRoute(static fs.FS) gin.HandlerFunc {
	httpFS := http.FS(static)

	return func(c *gin.Context) {
		if c.Request.URL.Path == "/config.json" {
			c.JSON(http.StatusOK, gin.H{
				"HEADER_NAME": os.Getenv("CRUDDY_HEADER_NAME"),
			})
			return
		}

		s, err := fs.Stat(static, path.Join("static"+c.Request.URL.Path))

		if c.Request.URL.Path == "/" || errors.Is(err, os.ErrNotExist) {
			c.FileFromFS("static/index.htm", httpFS)
			return
		}

		if s.IsDir() {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.FileFromFS("static"+c.Request.URL.Path, httpFS)
	}
}

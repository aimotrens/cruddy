package api

import (
	"os"
	"path"
)

type Handler struct {
	rootDir string
}

func NewHandler(rootDir string) *Handler {
	return &Handler{rootDir: rootDir}
}

func (h *Handler) getFullpathChecked(reqPath string) (string, bool) {
	fullpath := path.Join(h.rootDir, reqPath)

	_, err := os.Stat(fullpath)
	if os.IsNotExist(err) {
		return fullpath, false
	}

	return fullpath, true
}

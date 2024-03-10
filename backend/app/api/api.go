package api

import (
	"os"
	"path/filepath"
	"strings"
)

type Handler struct {
	rootDir string
}

func NewHandler(rootDir string) *Handler {
	return &Handler{rootDir: rootDir}
}

func (h *Handler) getFullpathChecked(reqPath string) (string, bool) {
	joinedReqPath := filepath.Join(h.rootDir, reqPath)

	rootFullpath, _ := filepath.Abs(h.rootDir)
	reqFullpath, _ := filepath.Abs(joinedReqPath)

	if !strings.HasPrefix(reqFullpath, rootFullpath) {
		return "", false
	}

	_, err := os.Stat(reqFullpath)
	if os.IsNotExist(err) {
		return "", false
	}

	return reqFullpath, true
}

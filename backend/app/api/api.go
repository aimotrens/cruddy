package api

import (
	"os"
	"path"
)

type Handler struct {
	base string
}

func NewHandler(base string) *Handler {
	return &Handler{base: base}
}

func (h *Handler) getFullpathChecked(reqPath string) (string, bool) {
	fullpath := path.Join(h.base, reqPath)

	_, err := os.Stat(fullpath)
	if os.IsNotExist(err) {
		return fullpath, false
	}

	return fullpath, true
}

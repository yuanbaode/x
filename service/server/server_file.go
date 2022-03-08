package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func serverFile(path string) gin.HandlerFunc {
	h := http.FileServer(gin.Dir("./", true))
	if path != "" {
		h = http.StripPrefix(path, h)
	}
	return gin.WrapH(h)
}

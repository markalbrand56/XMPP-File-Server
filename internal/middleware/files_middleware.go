package middlewares

import (
	"XMPP-File-Server/internal/configs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BodySizeLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, configs.MAX_UPLOAD_SIZE)
	}
}

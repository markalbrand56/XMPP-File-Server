package routes

import (
	"XMPP-File-Server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Service is up and running!",
		})
	})

	r.POST("/:directory", controllers.Upload)
	r.GET("/:directory/:file", controllers.GetFile)
}

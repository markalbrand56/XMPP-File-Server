package routes

import (
	"XMPP-File-Server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":       "File Service is up and running!",
			"documentation": "https://redes-markalbrand56.koyeb.app/files/documentation/index.html",
		})
	})

	r.POST("/:directory", controllers.Upload)
	r.GET("/:directory/:file", controllers.GetFile)
}

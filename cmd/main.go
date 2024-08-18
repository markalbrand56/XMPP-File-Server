package main

import (
	"XMPP-File-Server/docs"
	_ "XMPP-File-Server/internal/configs"
	"XMPP-File-Server/internal/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORS() gin.HandlerFunc {
	// Reference: https://github.com/gin-contrib/cors/issues/29
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(CORS())

	// Swagger
	docs.SwaggerInfo.Title = "Nomade Global Talent Website Backend"
	docs.SwaggerInfo.Description = "API para la el backend de la página web de Nomade Global Talent"
	docs.SwaggerInfo.Version = "0.1.0"
	docs.SwaggerInfo.Host = "ngt-markalbrand56.koyeb.app"
	docs.SwaggerInfo.BasePath = "/"

	// Routes
	//urlSwagger := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.Routes(r)

	err := r.Run() // listen and serve on
	if err != nil {
		panic(err)
	}
}

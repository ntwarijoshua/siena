package http

import (
	"github.com/gin-gonic/gin"
	"github.com/ntwarijoshua/siena/internal/http/Handlers"
)

func GetRouter(app Handlers.App) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{

			v1.POST("/auth", app.AuthenticateUser)
			v1.POST("/users", app.CreateUser)
			// protected end points
			protected := v1.Group("")
			{
				protected.Use(app.AuthMiddleware())
				protected.GET("/", func(context *gin.Context) {
					context.JSON(200, "SIENA-API v1")
				})

			}

		}
	}

	return r
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
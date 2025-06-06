package routes

import (
	"github.com/gin-gonic/gin"
	"007-go-api-redis/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/cache/:key", controllers.GetCache)
	r.POST("/cache/:key", controllers.SetCache)

	return r
}

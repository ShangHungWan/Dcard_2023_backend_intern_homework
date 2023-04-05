package router

import (
	"key-value-system/controllers"

	"github.com/gin-gonic/gin"
)

func Run(listen string) {
	router := gin.Default()

	registerRouter(router)

	router.Run(listen)
}

func registerRouter(router *gin.Engine) {
	router.POST("/head", controllers.StoreHead)
	router.POST("/page", controllers.StoreNode)
	router.GET("/head/:key", controllers.ShowHead)
	router.GET("/page/:key", controllers.ShowNode)
	router.DELETE("/head/:key", controllers.RemoveHead)
	router.DELETE("/page/:key", controllers.RemoveNode)
}

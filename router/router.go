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
	router.POST("/node", controllers.StoreNode)
	router.GET("/head/:key", controllers.ShowHead)
	router.GET("/node/:key", controllers.ShowNode)
	router.DELETE("/head/:key", controllers.RemoveHead)
	router.DELETE("/node/:key", controllers.RemoveNode)
}

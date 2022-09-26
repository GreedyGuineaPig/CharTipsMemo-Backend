package router

import (
	"gomysql-api/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.GetAll)
	// router.GET("/", controller.GetFirst)
	router.POST("/", controller.AddPost)
	router.PATCH("/", controller.PatchPost)
	router.DELETE("/:id", controller.DeletePost)
	router.GET("/:char", controller.GetCharpost)
	return router
}

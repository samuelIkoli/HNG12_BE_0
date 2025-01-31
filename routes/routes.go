package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/samuelIkoli/HNG12_BE_0/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Deployed and Running"})
	})

	router.GET("/task", controller.Task)
	router.GET("/test", controller.Test)
	router.GET("/ping", controller.Ping)
	router.GET("/HNG", controller.HNG)
}
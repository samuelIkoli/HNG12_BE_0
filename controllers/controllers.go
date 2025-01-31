package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/samuelIkoli/HNG12_BE_0/entity"
)

func Test(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "test working with air hot reload and refactoring",
	})
}

func Task(ctx *gin.Context){
	now:= time.Now().UTC()
	result := entity.Response{
		Email: "ayibanimiikoli@gmail.com",
		Current_datetime: now.Format(time.RFC3339),
		Github_url: "https://github.com/samuelIkoli/HNG12_BE_0",
	}
	ctx.JSON(200, result)
}

func Ping(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "Pong",
	})
}

func HNG(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"message": "HNG 12 - 2025, Mark Essien rocks :-)",
	})
}

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"github.com/samuelIkoli/HNG12_BE_0/entity"
)


func main(){

	c := cron.New()
	c.AddFunc("*/10 * * * *", func() {
		fmt.Println("Cronning")
		http.Get("https://hng12-be-0.onrender.com/")

	})
	c.Start()
	server:= gin.Default()

	server.GET("/", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Deployed and Running",
		})
	}) 

	server.GET("/task", func(ctx *gin.Context){
		now:= time.Now().UTC()
		result := entity.Response{
			Email: "ayibanimiikoli@gmail.com",
			Current_datetime: now.Format(time.RFC3339),
			Github_url: "https://github.com/samuelIkoli/HNG12_BE_0",
		}
		ctx.JSON(200, result)
	})

	server.GET("/test", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "test working with air hot reload",
		})
	})

	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	server.GET("/HNG", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "HNG 12 - 2025, Mark Essien rocks :-)",
		})
	})

	server.Run(":8080")

}
package main

import (
	"fmt",

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("checking server")
	r:=gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run()
}
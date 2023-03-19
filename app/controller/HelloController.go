package controller

import "github.com/gin-gonic/gin"

func Hello(ctx *gin.Context) {
	println("helloController invoke")
}

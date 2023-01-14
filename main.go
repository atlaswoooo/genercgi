package main

import (
	"genercgi/handler/login"
	"genercgi/handler/register"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/login", func(context *gin.Context) {
		login.Login(context)
	})
	r.POST("/register", func(context *gin.Context) {
		register.Register(context)
	})
	r.Run(":2000")
}

// 访问测试  http://54.187.249.98:2000

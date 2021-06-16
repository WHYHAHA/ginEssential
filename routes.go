package main

import (
	"github.com/gin-gonic/gin"
	"why/ginessential/controller"
	"why/ginessential/middle"
)

func CollectRoute(r * gin.Engine) * gin.Engine  {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info",middle.AuthMiddleware(), controller.Info)

	return  r
}
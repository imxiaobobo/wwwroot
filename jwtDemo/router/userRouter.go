package router

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/controller"
	"jwtDemo/middleware"
)

func UserRouter(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register) //注册
	r.POST("/login", controller.Login)
	r.GET("/info", middleware.Auth(), controller.Info)
	return r
}

package router

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/controller"
)

func UserRouter(r *gin.Engine) *gin.Engine {
	r.POST("/register", controller.Register) //注册
	r.POST("login", controller.Login)
	return r
}

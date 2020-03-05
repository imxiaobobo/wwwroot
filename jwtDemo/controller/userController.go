package controller

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/common"
	"jwtDemo/model"
	"jwtDemo/utils"
	"net/http"
)

/**
  注册
*/
func Register(ctx *gin.Context) {
	//绑定参数
	User := &model.UserModel{}
	ctx.ShouldBind(User)
	User.Db = common.NewDB()
	//数据验证
	if len(User.Phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if User.IsPhoneExist() { //判断手机号是否存在
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}
	if len(User.Password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须为6位"})
		return
	}
	if len(User.Name) == 0 { //随机给10位字符串做名字
		User.Name = utils.RandomName()
	}

	User.CreateUser()                             //创建用户
	ctx.JSON(http.StatusOK, gin.H{"msg": "注册成功"}) //返回结果
}

/**
  登录
*/
func Login(ctx *gin.Context) {
	//获取参数
	//数据验证
	//判断手机号是否存在
	//判断密码是否正确
	//发放token
}

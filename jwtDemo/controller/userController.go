package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost) //密码加密
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "密码加密错误"})
		return
	}
	User.Password = string(hashPassword)
	User.CreateUser()                             //创建用户
	ctx.JSON(http.StatusOK, gin.H{"msg": "注册成功"}) //返回结果
}

/**
  登录
*/
func Login(ctx *gin.Context) {
	//获取参数
	User := &model.UserModel{}
	ctx.ShouldBind(User)
	User.Db = common.NewDB()
	//数据验证
	if len(User.Phone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(User.Password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须为6位"})
		return
	}
	password := User.Password //下面判断手机号的时候已经将User重新赋值了,所以得先保存password,以供下面判断
	//判断手机号是否存在
	if !User.IsPhoneExist() { //判断手机号是否存在
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码错误"})
		return
	}
	//发放token
	token, err := common.ReleaseToken(*User)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "token生成失败" + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "token": token})
}

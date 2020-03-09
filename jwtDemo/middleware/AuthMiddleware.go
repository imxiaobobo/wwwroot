package middleware

import (
	"github.com/gin-gonic/gin"
	"jwtDemo/common"
	"jwtDemo/model"
	"net/http"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//jwt一般放在head里,取出来
		tokenStr := c.GetHeader("Authorization")
		//验证token
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		tokenStr = tokenStr[7:] //去掉头部的Bearer
		token, claims, err := common.ParseToken(tokenStr)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "token验证不通过"})
			c.Abort()
			return
		}
		userId := claims.UserId
		Db := common.NewDB()
		var user model.UserModel
		Db.First(&user, userId)
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

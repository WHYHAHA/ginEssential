package middle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"why/ginessential/common"
	"why/ginessential/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")

		// validate token formate  author2.0规定字符串开头必须要有Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString,"Bearer") {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]

		token,claims,err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}

		// 验证通过后获取claim 中的userId
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user,userId)
		
		//用户 不存在
		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized,gin.H{"code":401,"msg":"权限不足"})
			ctx.Abort()
			return
		}

		//用户 存在 将user信息写入上下文
		ctx.Set("user",user)

		ctx.Next()
	}
}
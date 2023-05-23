package middleware

import (
	"gin-mall/pkg/e"
	"gin-mall/pkg/util"
	"github.com/gin-gonic/gin"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		code = 200
		token := context.GetHeader("authorization")
		if token == "" {
			code = e.Error
		} else { //token 解析和过期
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ErrorAuthToken
			} else if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
				code = e.ErrorAuthCheckTokenTimeOut
			}
		}

		if code != 200 {
			context.JSON(code, gin.H{
				"status": code,
				"Msg":    e.GetMsg(code),
			})
			context.Abort()
			return
		}

		context.Next()

	}

}

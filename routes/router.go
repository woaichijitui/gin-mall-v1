package routes

import (
	"gin-mall/api"
	"github.com/gin-gonic/gin"
)
import (
	"gin-mall/middleware"
)

func NewRouter() *gin.Engine {
	route := gin.Default()

	route.Use(middleware.Cors())
	route.Static("/static", "./static") //加载静态路径

	//路由分组
	v1 := route.Group("api/v1")

	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "sussuce")
		})

		v1.POST("user/register", api.UserRegister)
	}
	return route
}

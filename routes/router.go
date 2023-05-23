package routes

import (
	api "gin-mall/api/v1"
	"gin-mall/middleware"
	"github.com/gin-gonic/gin"
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
		}) //测试

		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)

		authed := v1.Group("/") //登录保护
		{
			authed.PUT("/user", api.UserUpdate)
			authed.POST("/avatar", api.UpdateAvatar)
		}
	}
	return route
}

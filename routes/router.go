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

		//商品操作
		v1.GET("/products", api.ListProduct)
		v1.GET("/products/:id", api.ShowProduct)

		//轮播图
		v1.POST("/carousel", api.ListCarousel)
		authed := v1.Group("/") //登录保护
		{
			authed.PUT("/user", api.UserUpdate)
			authed.POST("/avatar", api.UpdateAvatar)
			authed.POST("/user/sending-email", api.SendEmail)
			authed.POST("/user/valid-email", api.ValidEmail)
			authed.POST("/money", api.ShowMoney)

			//显示金额
			authed.POST("/")

			//创建商品
			authed.POST("/product", api.CreateProduct)
		}
	}
	return route
}

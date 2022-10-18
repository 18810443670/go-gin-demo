package Router

import (
	"gin-1/Controllers"
	"gin-1/Middlewares"
	"github.com/gin-gonic/gin"
)

// 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//程序运行时间
	r.Use(Middlewares.MyTime)

	//跨域
	r.Use(Middlewares.Cors())

	//登录
	login := r.Group("/v1")
	{
		con := &Controllers.UserController{}
		login.POST("login", con.Login)
		login.POST("register", con.Register)
	}

	//jwt授权接口
	jwt := r.Group("/v1").Use(Middlewares.JWTAuthMiddleware)
	{
		con := &Controllers.UserController{}
		jwt.GET("getInfo", con.GetInfo)

		//创建标签
		tag := &Controllers.TagController{}
		jwt.POST("tag/create", tag.CreateTags)

	}

	return r
}

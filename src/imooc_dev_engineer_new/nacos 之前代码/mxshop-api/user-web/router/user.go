package router

import (
	"mxshop-api/user-web/api"
	"mxshop-api/user-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user") // .Use(middlewares.JWTAuth())
	{
		UserRouter.GET("list",
			middlewares.JWTAuth(),
			middlewares.IsAdminAuth(),
			api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		// 注册接口
		UserRouter.POST("register", api.Register)
	}
}

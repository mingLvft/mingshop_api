package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mingshop_api/user_web/api"
	"mingshop_api/user_web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("/user")
	zap.S().Info("配置用户路由")
	{
		UserRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("/pwd_login", api.PassWordLogin)
	}
}
